package controllers

import (
	"errors"
	"net/http"
	"sea-study/api/models"
	"sea-study/constants"
	"sea-study/service"
	"sea-study/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(c *gin.Context, db *gorm.DB) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidInput})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrFailedToHash})
        return
    }
    user.Password = string(hashedPassword)

    // Call the service layer to create the user
    if err := service.CreateUser(db, &user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrFailedToCreateUser})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": user})
}

func LoginUser(c *gin.Context, db *gorm.DB) {
	var loginCredentials struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&loginCredentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidInput})
		return
	}

	var user models.User
	if err := service.GetUserByEmail(db, &user, loginCredentials.Email); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrInvalidCredentials})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrFailedToRetrieveUser})
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginCredentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrInvalidCredentials})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserID,
		"role": user.Role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(util.GetJWTSecret()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrFailedToGenerateToken})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
		"user": gin.H{
			"id":    user.UserID,
			"name": user.Name,
			"email": user.Email,
			"role": user.Role,
		},
	})
}

func GetUserProfile(c *gin.Context, db *gorm.DB) {
	// Get the user ID from the context (set by the UserMiddleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ErrUserNotAuthenticated})
		return
	}

	// Convert the userID to UUID
	userUUID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrInvalidUserID})
		return
	}

	var user models.User
	if err := service.GetUserByID(db, &user, userUUID); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": constants.ErrUserNotFound})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrFailedToRetrieveUserProfile})
		}
		return
	}

	response := gin.H{
		"userId":    user.UserID,
		"name":      user.Name,
		"email":     user.Email,
		"balance":   user.Balance,
		"role":      user.Role,
		"createdAt": user.CreatedAt,
		"updatedAt": user.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}
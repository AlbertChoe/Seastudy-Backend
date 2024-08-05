package controllers

import (
	"net/http"
	"sea-study/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TopupInput struct {
	Amount        float64 `json:"amount" binding:"required"`
	PaymentMethod string  `json:"payment_method" binding:"required"`
}

func Topup(c *gin.Context, db *gorm.DB) {
	var input TopupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	topup, err := service.CreateTopup(db, userID.(string), input.Amount, input.PaymentMethod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Top-up successful", "topup": topup})
}

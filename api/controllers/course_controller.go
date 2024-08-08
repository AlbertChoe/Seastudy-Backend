package controllers

import (
	"fmt"
	"net/http"
	"os"
	"sea-study/api/models"
	"sea-study/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateCourse(c *gin.Context, db* gorm.DB) {
	var input models.CourseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	course, err := service.CreateCourse(db, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Course created successfully", "course": course})
}

func GetAllCourses(c *gin.Context, db *gorm.DB) {
    courses, err := service.GetAllCourses(db)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve courses"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"courses": courses})
}

func GetCourse(c *gin.Context, db *gorm.DB) {
    courseIDParam := c.Param("course_id")
    courseID, err := strconv.Atoi(courseIDParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }
    
    course, err := service.GetCourse(db, courseID)
    if err != nil {
        if err.Error() == "course not found" {
            c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve course"})
        }
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"course": course})
}

func UpdateCourse(c *gin.Context, db *gorm.DB) {
    courseIDParam := c.Param("course_id")
    courseID, err := strconv.Atoi(courseIDParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    var input models.CourseInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    course, err := service.UpdateCourse(db, courseID, &input)
    if err != nil {
        if err.Error() == "course not found" {
            c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course"})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully", "course": course})
}

func DeleteCourse(c *gin.Context, db *gorm.DB) {
    courseIDParam := c.Param("course_id")
    courseID, err := strconv.Atoi(courseIDParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    err = service.DeleteCourse(db, courseID)
    if err != nil {
        if err.Error() == "course not found" {
            c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course"})
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}

func UploadCourseImage(c *gin.Context, db *gorm.DB) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload image"})
		return
	}

	extension := file.Filename[len(file.Filename)-4:]

	imageID := uuid.New().String()

	filePath := fmt.Sprintf("uploads/%s%s", imageID, extension)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

    hostURL := os.Getenv("HOST_URL")
	imageURL := fmt.Sprintf("%s/%s", hostURL, filePath)

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "image_url": imageURL})
}
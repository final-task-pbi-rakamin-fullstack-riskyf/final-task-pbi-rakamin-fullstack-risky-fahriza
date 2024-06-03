package controllers

import (
	"myapp/database"
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	var input models.Photo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}
	input.UserID = userId.(uint)

	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	if err := database.DB.Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": photos})
}

func UpdatePhoto(c *gin.Context) {
	var photo models.Photo
	id := c.Param("photoId")
	if err := database.DB.Where("id = ?", id).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	var input models.Photo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&photo).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": photo})
}

func DeletePhoto(c *gin.Context) {
	var photo models.Photo
	id := c.Param("photoId")
	if err := database.DB.Where("id = ?", id).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	database.DB.Delete(&photo)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
package controllers

import (
	"fmt"
	"gostud/initializers"
	"gostud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Success login",
		"data":    user,
	})
}

func UpdateUser(c *gin.Context) {
	type requestBody struct {
		Name  string
		Email string
	}
	var json requestBody
	err := c.ShouldBindJSON(&json)
	fmt.Println("Debugging: ", json.Name == "")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "data must be json",
		})
		return
	}
	if json.Email == "" || json.Name == "" {
		var message []string
		if json.Email == "" {
			message = append(message, "data email is required")
		}

		if json.Name == "" {
			message = append(message, "data name is required")
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"errors":  message,
			"message": "failed request body",
		})
		return
	}

	userId, _ := c.Get("userId")
	var modelUser models.User
	initializers.DB.Where("id = ?", userId).First(&modelUser)
	modelUser.Name = json.Name
	modelUser.Email = json.Email
	initializers.DB.Save(&modelUser)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Success to updated data user",
		"data":    modelUser,
	})
}

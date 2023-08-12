package controllers

import (
	"gostud/initializers"
	"gostud/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":  "Azis Zuhri Pratomo",
		"kelas": "1 D4 IT A",
	})
}

func SignUp(c *gin.Context) {
	var requestBody models.UserModel
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid request body",
			"errors":  err.Error(),
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "failed to hash password",
		})
		return
	}
	var checkUser models.User
	err = initializers.DB.Where("email = ?", requestBody.Email).First(&checkUser).Error
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "user was created",
		})

		return
	}

	// Create User
	user := models.User{
		Name:      requestBody.Name,
		Email:     requestBody.Email,
		Password:  string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "failed to create user",
			"errors":  result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Success to create user",
	})
}

func Login(c *gin.Context) {
	var requestBody struct {
		Email    string
		Password string
	}
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "email and password required",
		})
		return
	}

	var checkUser models.User
	err = initializers.DB.Where("email = ?", requestBody.Email).First(&checkUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "email or password was wrong",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(requestBody.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "email or password was wrong",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": checkUser.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(viper.Get("JWT_SECRET").(string)))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "failed to create token",
			"error":   err.Error(),
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"data":    checkUser,
		"token":   tokenString,
		"message": "success to created token",
	})
}

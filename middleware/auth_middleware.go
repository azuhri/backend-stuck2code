package middleware

import (
	"fmt"
	"gostud/initializers"
	"gostud/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func Auth(c *gin.Context) {
	fmt.Println("========= IN MIDDLEWARE AUTH =============")

	tokenstring, err := c.Cookie("Auth")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenstring, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected sigining method: %v", t.Header["alg"])
		}

		return []byte(viper.Get("JWT_SECRET").(string)), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		err = initializers.DB.First(&user, claims["sub"]).Error
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)
		c.Set("userId", claims["sub"])
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}

package routes

import (
	"gostud/controllers"
	"gostud/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRoute(route *gin.Engine) {
	v1 := route.Group("/api/v1")

	auth := v1.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/signup", controllers.SignUp)
	}

	app := v1.Group("/app")
	app.Use(middleware.Auth)
	{
		// USER ROUTES
		userRoutes := app.Group("/user")
		{
			userRoutes.GET("/", controllers.GetUser)
			userRoutes.PUT("/", controllers.UpdateUser)
		}

		// // CHANNEL ROUTES
		// channelRoutes := app.Group("/channel")
		// {
		// 	channelRoutes.POST("/", controllers.CreateChannel)

		// 	memberChannelRoutes := channelRoutes.Group("/member")
		// 	{
		// 		memberChannelRoutes.POST("/", controllers.AddMember)
		// 		memberChannelRoutes.DELETE("/", controllers.DeleteMember)
		// 	}
		// }
	}
}

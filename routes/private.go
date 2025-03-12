package routes

import (
	handlers "blog/handlers/private"

	"github.com/gin-gonic/gin"
)

func SetupPrivateRoutes(router *gin.Engine) {
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// Comments
	commentsRoutes := router.Group("/comments")
	{
		commentsRoutes.GET("/", handlers.GetAllComments)
		commentsRoutes.GET("/:id", handlers.GetComment)
		commentsRoutes.POST("/", handlers.CreateComment)
		commentsRoutes.PUT("/:id", handlers.UpdateComment)
		commentsRoutes.DELETE("/:id", handlers.DeleteComment)
	}

	// Users
	usersRoutes := router.Group("/users")
	{
		usersRoutes.GET("/", handlers.GetAllUsers)
		usersRoutes.GET("/:id", handlers.GetUser)
		usersRoutes.POST("/", handlers.CreateUser)
		usersRoutes.PUT("/:id", handlers.UpdateUser)
		usersRoutes.DELETE("/:id", handlers.DeleteUser)
	}
}

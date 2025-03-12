package routes

import (
	handlers "blog/handlers/private"
	"blog/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	setupPublicRoutes(router)
	setupProtectedUserRoutes(router)
	setupProtectedCommentRoutes(router)
}

func setupPublicRoutes(router *gin.Engine) {
	// public := router.Group("/api/v1/public")
	// {
	// 	public.GET("/me", handlers.Me)
	// 	public.GET("/health", handlers.HealthHandler)
	// 	public.POST("/login", handlers.Login)
	// 	public.POST("/signup", handlers.Signup)
	// }
}

func setupProtectedUserRoutes(router *gin.Engine) {
	usersRoutes := router.Group("/api/v1/users")
	usersRoutes.Use(middlewares.AuthMiddleware())
	{
		usersRoutes.GET("/", handlers.GetAllUsers)
		usersRoutes.GET("/:id", middlewares.ParseId(), handlers.GetUser)
		usersRoutes.POST("/", middlewares.ParseId(), handlers.CreateUser)
		usersRoutes.PUT("/:id", middlewares.ParseId(), handlers.UpdateUser)
		usersRoutes.DELETE("/:id", middlewares.ParseId(), handlers.DeleteUser)
	}
}

func setupProtectedCommentRoutes(router *gin.Engine) {
	commentsRoutes := router.Group("/api/v1/comments")
	commentsRoutes.Use(middlewares.AuthMiddleware())
	{
		commentsRoutes.GET("/", handlers.GetAllComments)
		commentsRoutes.GET("/:id", middlewares.ParseId(), handlers.GetComment)
		commentsRoutes.POST("/", handlers.CreateComment)
		commentsRoutes.PUT("/:id", middlewares.ParseId(), handlers.UpdateComment)
		commentsRoutes.DELETE("/:id", middlewares.ParseId(), handlers.DeleteComment)
	}
}

// func setupProtectedPostRoutes(router *gin.Engine) {
// 	protected := router.Group("/api/v1/post")
// 	protected.Use(middlewares.AuthMiddleware())
// 	{
// 		protected.GET("/", handlers.GetPosts)
// 		protected.GET("/:id", handlers.GetPost)
// 		protected.POST("/", handlers.CreatePost)
// 		protected.PUT("/:id", handlers.UpdatePost)
// 		protected.DELETE("/:id", handlers.DeletePost)

// 		protected.POST("/:id/like", handlers.LikePost)
// 		protected.POST("/:id/unlike", handlers.UnlikePost)
// 	}
// }

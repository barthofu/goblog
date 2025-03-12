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
	setupProtectedPostRoutes(router)
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

func setupProtectedPostRoutes(router *gin.Engine) {
	postsRoutes := router.Group("/api/v1/posts")
	postsRoutes.Use(middlewares.AuthMiddleware())
	{
		postsRoutes.GET("/", handlers.GetAllPosts)
		postsRoutes.GET("/:id", middlewares.ParseId(), handlers.GetPost)
		postsRoutes.POST("/", handlers.CreatePost)
		postsRoutes.PUT("/:id", middlewares.ParseId(), handlers.UpdatePost)
		postsRoutes.DELETE("/:id", middlewares.ParseId(), handlers.DeletePost)

		postsRoutes.POST("/:id/like", middlewares.ParseId(), handlers.LikePost)
		postsRoutes.POST("/:id/unlike", middlewares.ParseId(), handlers.UnlikePost)
	}
}

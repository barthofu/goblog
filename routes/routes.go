package routes

import (
	"blog/handlers"
	"blog/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	setupPublicRoutes(router)
	setupProtectedUserRoutes(router)
	setupProtectedCommentRoutes(router)
	setupProtectedArticleRoutes(router)
}

func setupPublicRoutes(router *gin.Engine) {
	public := router.Group("/api/v1/public")
	{
		public.GET("/health", handlers.HealthCheck)
		public.POST("/register", handlers.Register)
	}
}

func setupProtectedUserRoutes(router *gin.Engine) {
	usersRoutes := router.Group("/api/v1/users")
	usersRoutes.Use(middlewares.AuthMiddleware())
	{
		usersRoutes.GET("/", handlers.GetAllUsers)
		usersRoutes.GET("/:id", middlewares.ParseId(), handlers.GetUser)
		usersRoutes.POST("/", handlers.CreateUser)
		usersRoutes.PUT("/:id", middlewares.ParseId(), handlers.UpdateUser)
		usersRoutes.DELETE("/:id", middlewares.ParseId(), handlers.DeleteUser)

		usersRoutes.GET("/me", handlers.Me)
		usersRoutes.POST("/:id/follow", middlewares.ParseId(), handlers.FollowUser)
		usersRoutes.POST("/:id/unfollow", middlewares.ParseId(), handlers.UnfollowUser)
	}
}

func setupProtectedCommentRoutes(router *gin.Engine) {
	commentsRoutes := router.Group("/api/v1/comments")
	commentsRoutes.Use(middlewares.AuthMiddleware())
	{
	}
}

func setupProtectedArticleRoutes(router *gin.Engine) {
	articlesRoutes := router.Group("/api/v1/articles")
	articlesRoutes.Use(middlewares.AuthMiddleware())
	{
		articlesRoutes.GET("/", handlers.GetAllArticles)
		articlesRoutes.GET("/:id", middlewares.ParseId(), handlers.GetArticle)
		articlesRoutes.POST("/", handlers.CreateArticle)
		articlesRoutes.PUT("/:id", middlewares.ParseId(), handlers.UpdateArticle)
		articlesRoutes.DELETE("/:id", middlewares.ParseId(), handlers.DeleteArticle)
		articlesRoutes.POST("/:id/comment", middlewares.ParseId(), handlers.CreateComment)

		articlesRoutes.POST("/:id/like", middlewares.ParseId(), handlers.LikeArticle)
		articlesRoutes.POST("/:id/unlike", middlewares.ParseId(), handlers.UnlikeArticle)
	}
}

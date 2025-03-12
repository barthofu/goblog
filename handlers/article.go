package handlers

import (
	"blog/models"
	"blog/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ==========================================
// CRUD Handlers
// ==========================================

func GetArticle(c *gin.Context) {
	id := c.GetInt("id")

	var article, err = services.GetArticle(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, article)
}

func GetAllArticles(c *gin.Context) {
	var articles, err = services.GetAllArticles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, articles)
}

func CreateArticle(c *gin.Context) {
	var input models.CreateOrUpdateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user = c.MustGet("user").(*models.User)

	article, err := services.CreateArticle(input.Title, input.Content, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, article)
}

func UpdateArticle(c *gin.Context) {
	var id = c.GetInt("id")
	var input models.CreateOrUpdateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := services.UpdateArticle(id, input.Title, input.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, article)
}

func DeleteArticle(c *gin.Context) {
	var id = c.GetInt("id")

	err := services.DeleteArticle(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// ==========================================
// Likes Handlers
// ==========================================

func LikeArticle(c *gin.Context) {
	var id = c.GetInt("id")
	var user = c.MustGet("user").(*models.User)

	err := services.LikeArticle(id, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func UnlikeArticle(c *gin.Context) {
	var id = c.GetInt("id")
	var user = c.MustGet("user").(*models.User)

	err := services.UnlikeArticle(id, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

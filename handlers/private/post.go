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

func GetPost(c *gin.Context) {
	id := c.GetInt("id")

	var post, err = services.GetPost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func GetAllPosts(c *gin.Context) {
	var posts, err = services.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
	var input models.CreateOrUpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user = c.MustGet("user").(*models.User)

	post, err := services.CreatePost(input.Title, input.Content, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func UpdatePost(c *gin.Context) {
	var id = c.GetInt("id")
	var input models.CreateOrUpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := services.UpdatePost(id, input.Title, input.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func DeletePost(c *gin.Context) {
	var id = c.GetInt("id")

	err := services.DeletePost(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// ==========================================
// Likes Handlers
// ==========================================

func LikePost(c *gin.Context) {
	var id = c.GetInt("id")
	var user = c.MustGet("user").(*models.User)

	err := services.LikePost(id, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func UnlikePost(c *gin.Context) {
	var id = c.GetInt("id")
	var user = c.MustGet("user").(*models.User)

	err := services.UnlikePost(id, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

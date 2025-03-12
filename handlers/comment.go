package handlers

import (
	"blog/models"
	"blog/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var articleId = c.GetInt("id")
	var input models.CreateCommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user = c.MustGet("user").(*models.User)

	comment, err := services.CreateComment(input.Content, articleId, *user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

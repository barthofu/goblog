package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseId() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid `id`"})
			c.Abort()
			return
		}
		c.Set("id", id) // Stocker l'ID dans le contexte
		c.Next()
	}
}

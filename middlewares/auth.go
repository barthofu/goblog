package middlewares

import (
	"blog/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.GetHeader("X-User-Email")
		if email == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing email header"})
			c.Abort()
			return
		}

		// Vérification en base de données
		var user, err = services.GetUserByEmail(email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Stocke l'utilisateur dans le contexte Gin
		c.Set("user", user)

		// Continue l'exécution des handlers
		c.Next()
	}
}

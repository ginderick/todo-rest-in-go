package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	userss := []User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
	}

	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"users": userss,
		})
	})
}

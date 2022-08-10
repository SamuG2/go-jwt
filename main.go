package main

import (
	"os"

	routes "github.com/SamuG2/go-jwt/routes"

	"github.com/gin-gonic/gin"
)

// @title JWT API
// @version 1.0
// @description  Api para registro de usuarios con go

// @contact.name S
// @contact.email prueba@gmail.com

// @license.name lisensiado
// @license.url lisensiado.com

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		routes.AuthRoutes(users)
		routes.UserRoutes(users)
	}

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}

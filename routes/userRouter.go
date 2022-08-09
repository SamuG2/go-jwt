package routes

import (
	controller "github.com/SamuG2/go-jwt/controllers"
	"github.com/SamuG2/go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:users_id", controller.GetUser())
}

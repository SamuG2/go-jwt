package routes

import (
	controller "github.com/SamuG2/go-jwt/controllers"
	"github.com/SamuG2/go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("", controller.GetUsers())
	incomingRoutes.GET(":users_id", controller.GetUser())
}

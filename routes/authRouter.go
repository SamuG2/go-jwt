package routes

import (
	controller "github.com/SamuG2/go-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.POST("signup", controller.SignUp())
	incomingRoutes.POST("login", controller.Login())
}

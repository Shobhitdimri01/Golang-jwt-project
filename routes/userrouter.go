package routes

import (
	"github.com/Shobhitdimri01/jwt-project/middleware"

	controller "github.com/Shobhitdimri01/jwt-project/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}

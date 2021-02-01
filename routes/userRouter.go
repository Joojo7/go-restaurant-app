package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/heroku/go-getting-started/controllers"
)

//UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/signup", controller.CreateUser())
	incomingRoutes.POST("/users/login", controller.Login())
}
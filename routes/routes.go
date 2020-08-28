package routes

import (
	"go_web_app/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
	return router
}

package routes

import (
	"go_web_app/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", controllers.CreateUser)
			users.GET("", controllers.GetUsers)
			users.GET(":id", controllers.GetUser)
			users.DELETE(":id", controllers.DeleteUser)
		}
	}
	return router
}

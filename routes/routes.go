package routes

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	router := gin.Default()
	router.GET("/user", users.getUsers)

}

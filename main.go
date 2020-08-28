package main

import (
	"go_web_app/gostrings"
	"go_web_app/users"

	"github.com/gin-gonic/gin"
)

func main() {
	gostrings.Squish("sssssss")
	r := gin.Default()
	r.GET("/users", users.GetUsers)
	r.Run() // listen and serve on 0.0.0.0:8080
}

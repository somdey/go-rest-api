package main

import (
	"go_web_app/routes"
)

func main() {
	router := routes.InitializeRoutes()
	router.Run() // listen and serve on 0.0.0.0:8080
}

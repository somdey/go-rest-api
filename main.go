package main

import (
	"context"
	"fmt"
	_ "go_web_app/docs"
	"go_web_app/routes"
	"log"
	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title Golang api doc
// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://somdey:" + os.Getenv("dbPass") + "@ds125041.mlab.com:25041/go_web_app")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	router := routes.InitializeRoutes()
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") //The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}

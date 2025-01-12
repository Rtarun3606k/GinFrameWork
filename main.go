package main

import (
	"context"
	"fmt"
	"log"

	routes "GinFrameWork/Routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Create a context
	ctx := context.TODO()

	router := gin.Default()

	// Initialize UserController and set up routes
	userController := routes.NewUserController(client)
	userController.BasicRoute(router, ctx)

	router.Run(":8080")
}

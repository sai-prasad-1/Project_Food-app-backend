package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gin-gonic/gin"
	config "github.com/sai-prasad-1/Project_Food-app-backend/config"
	logger "github.com/sai-prasad-1/Project_Food-app-backend/loggers"
)

var collection *mongo.Collection
var ctx = context.TODO()

func main() {
	config := config.InitConfig()
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%v:%v@cluster0.nyyrtr9.mongodb.net/?retryWrites=true&w=majority", config.Database.MongoUser, config.Database.MongoPassword))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("tasker").Collection("tasks")
	fmt.Println("Connected to MongoDB!")

	logger.Info.Println("Port is num\t\t", config.Server.Port)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(fmt.Sprintf(": %v", config.Server.Port)) // listen and serve on the port specified
}

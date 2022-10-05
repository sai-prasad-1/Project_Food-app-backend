package dbConnection

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	config "github.com/sai-prasad-1/Project_Food-app-backend/config"
)

var UserCollection *mongo.Collection
var ctx = context.TODO()

func Connect() {
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
	UserCollection = client.Database("Cluster0").Collection("User")

	fmt.Println("Connected to MongoDB!")
}

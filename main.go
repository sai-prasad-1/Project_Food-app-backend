package main

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
	config "github.com/sai-prasad-1/Project_Food-app-backend/config"
	controllers "github.com/sai-prasad-1/Project_Food-app-backend/controllers"
	db "github.com/sai-prasad-1/Project_Food-app-backend/dbConnection"
)

var Collection *mongo.Collection

// var ctx = context.TODO()

func main() {
	config := config.InitConfig()
	db.Connect()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": Collection.Database(),
		})
	})

	r.GET("/user", controllers.UserLogin)
	r.POST("/user", controllers.UserCreate)

	r.Run(fmt.Sprintf(": %v", config.Server.Port)) // listen and serve on the port specified
}

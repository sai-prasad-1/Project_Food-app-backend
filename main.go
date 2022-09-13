package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	config "github.com/sai-prasad-1/Project_Food-app-backend/config"
)

func main() {
	config := config.InitConfig()
	fmt.Println("Port is num\t\t", config.Server.Port)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf(": %v", config.Server.Port)) // listen and serve on
}

package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	userCollection "github.com/sai-prasad-1/Project_Food-app-backend/dbConnection"
	models "github.com/sai-prasad-1/Project_Food-app-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type EmailRequestBody struct {
	Email    string
	Password string
}

type CreateRequestBody struct {
	Name     string
	Email    string
	Password string
	Phone    string
}

func UserLogin(c *gin.Context) {
	var requestBody EmailRequestBody
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	userCollection.UserCollection.FindOne(context.TODO(), bson.D{{Key: "email", Value: requestBody.Email}}).Decode(&user)
	if user.Name == "" {
		c.JSON(404, gin.H{
			"message": "Invalid Email",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password))
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Invalid Password",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": user,
	})

}

func UserCreate(c *gin.Context) {
	type UniqueEmail struct {
		Email string
	}
	var emailCheck UniqueEmail
	var requestBody models.User
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userCollection.UserCollection.FindOne(context.TODO(), bson.D{{Key: "email", Value: requestBody.Email}}).Decode(&emailCheck)
	if emailCheck.Email != "" {
		c.JSON(400, gin.H{
			"message": "Email already exists",
		})
		return
	}
	pass := []byte(requestBody.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		Name:      requestBody.Name,
		Email:     requestBody.Email,
		Password:  string(hashedPassword),
		Phone:     requestBody.Phone,
		Verified:  false,
	}

	result, err := userCollection.UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": user,
			"ID":      result,
		})
		return
	}

}

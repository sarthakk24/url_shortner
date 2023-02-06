package authentication

import (
	"context"
	"fmt"
	"net/http"
	"urlShortner/pkg/database"
	errorHandler "urlShortner/pkg/errors"
	"urlShortner/pkg/models/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate *validator.Validate

type doc struct {
	name string
}

func SignUp(c *gin.Context) {
	validate = validator.New()
	var newUser user.User
	err := c.ShouldBind(&newUser)

	if err != nil {
		errorHandler.Handler(err, c)
		return
	}

	newUser.InitUser()

	var userCollection mongo.Collection = database.GetCollection("urlShortner", "users")
	result, err := userCollection.InsertOne(context.TODO(), newUser)
	fmt.Println(result)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "New user was registered :)", "mongoId": result})
	return
}

func SignIn(c *gin.Context) {
	// validate = validator.New()

	// var user user.User
	// var details user.Login

	// err := c.ShouldBind(&details)

	// if err != nil {
	// 	errorHandler.Handler(err, c)
	// 	return
	// }

	// var userCollection mongo.Collection = database.GetCollection("urlShortner", "users")

	// var foundUser user.User:= userCollection.FindOne()

	// hashing.ComparePassword()

	c.IndentedJSON(http.StatusOK, gin.H{"message": "hello sign up initiated "})
}

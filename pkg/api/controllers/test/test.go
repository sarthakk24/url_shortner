package test

import (
	"context"
	"net/http"
	"urlShortner/pkg/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func TestMongo(c *gin.Context) {
	userCollection := database.GetCollection("urlShortner", "users")

	cursor, err := userCollection.Find(context.TODO(), bson.D{})

	var results []bson.M

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, results)
	return
}

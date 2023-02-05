package database

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client mongo.Client

func Connect() mongo.Client {

	var uri string = viper.GetString("MONGO_URI")
	fmt.Println(uri)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		fmt.Println(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("new connection to mongo db")
	return *client
}

func GetCollection(database string, collection string) mongo.Collection {
	var client mongo.Client = Connect()
	registeredCollection := client.Database(database).Collection(collection)
	return *registeredCollection
}

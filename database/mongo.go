package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func NewConnection() (*mongo.Client, error)  {
	log.Println("Stating Database")

	mongoUrl := mongoConnectionString()
	var client *mongo.Client

	if mongoUrl == "" {
		log.Println("No Connection String Provided!")
	} else {
		var err error
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))

		if err != nil {
			log.Fatal(err)
		}
	}
	return client,nil
}

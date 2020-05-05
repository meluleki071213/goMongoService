package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"goMongoService/database"
	"log"
	"time"
)

func main()  {
	var db *mongo.Client
	var err error

	for true {
		db, err = database.NewConnection()
		if err != nil {
			fmt.Println("Failed to access MongoDB. Retrying in 60 seconds")
			time.Sleep(1 * time.Minute)
		} else {
			break
		}
	}
	err = db.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
}

package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"goMongoService/database"
	"time"
)

func main()  {
	var _ *mongo.Client
	var err error

	for true {
		_, err = database.NewConnection()
		if err != nil {
			fmt.Println("Failed to access MongoDB. Retrying in 60 seconds")
			time.Sleep(1 * time.Minute)
		} else {
			fmt.Println("Database Connected")
			break
		}
	}
}

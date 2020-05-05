package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"goMongoService/database"
	"goMongoService/handlers"
	"log"
	"net/http"
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

	//Init Router
	r := mux.NewRouter()

	// arrange our route
	r.HandleFunc("/api/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", handlers.GetBook).Methods("GET")
	r.HandleFunc("/api/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", handlers.DeleteBook).Methods("DELETE")

	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))
}

package main

import (
	"log"
	"net/http"

	"github.com/Diego-Paris/tea-service/pkg/config"
	"github.com/Diego-Paris/tea-service/pkg/routes"
)

func main() {

	// Attempt connection to the database
	err := config.SetupDB()
	if err != nil {
		log.Fatalln("Could not connect to database.\n", err)
	}
	log.Println("Connected to Database!")

	// Setup all routes for the application
	router := routes.NewRouter()

	// Serve and run application
	log.Fatal(http.ListenAndServe(config.Port, router))
}

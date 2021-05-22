package main

import (
	"log"
	"net/http"

	"github.com/Diego-Paris/tea-service/pkg/config"
	"github.com/Diego-Paris/tea-service/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {

	// Attempt connection to the database
	err := config.SetupDB()
	if err != nil {
		log.Fatalln("Could not connect to database.\n", err)
	}
	log.Println("Connected to Database!")

	// Setup all routes for the application
	router := mux.NewRouter().StrictSlash(true)
	routes.InitializeAllRoutes(router)

	// Serve and run application
	log.Fatal(http.ListenAndServe(":8080", router))
}

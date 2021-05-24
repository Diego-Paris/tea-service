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
	//router := mux.NewRouter().StrictSlash(true)
	//routes.InitializeAllRoutes(router)
	router := routes.NewNewRouter()

	// m := http.NewServeMux()
	// m.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	req.URL.Path = utils.AddTrailingSlash(req.URL.Path)
	// 	router.ServeHTTP(w, req)
	// })

	// Serve and run application
	log.Fatal(http.ListenAndServe(config.Port, router)) //! change m back to router
}

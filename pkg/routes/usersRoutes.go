package routes

import (
	"github.com/Diego-Paris/tea-service/pkg/controllers"
	"github.com/gorilla/mux"
)

// initializeUserRoutes will set up all routes that begin with
// users i.e. "/users", "/users/{id}" etc.
func initializeUserRoutes(prefix string, router *mux.Router) {

	subrouter := router.PathPrefix(prefix).Subrouter() // set up subrouter for prefix

	// Get all users
	subrouter.HandleFunc("/", controllers.GetAllUsers).Methods("GET")
	
	// Get user by ID
	subrouter.HandleFunc("/{id}", controllers.GetUserByID).Methods("GET")
}

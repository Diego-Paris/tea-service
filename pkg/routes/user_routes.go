package routes

import (
	"github.com/Diego-Paris/tea-service/pkg/controllers"
	"github.com/gorilla/mux"
)

// UserRoutes contains all routes for the
// "/users" endpoint.
var UserRoutes = Routes{
	Route{
		"GetAllUsers",
		"GET",
		"/",
		controllers.GetAllUsers,
	},
	Route{
		"CreateUser",
		"POST",
		"/",
		controllers.CreateUser,
	},
	Route{
		"GetImage",
		"GET",
		"/image",
		controllers.GetImage,
	},
	Route{
		"GetUserByID",
		"GET",
		"/{id}",
		controllers.GetUserByID,
	},
}

// initializeUserRoutes will set up all routes that begin with
// users i.e. "/users", "/users/{id}" etc.
func initializeUserRoutes(prefix string, router *mux.Router) {

	subrouter := router.PathPrefix(prefix).Subrouter() // set up subrouter for prefix

	// Get all users
	subrouter.HandleFunc("/", controllers.GetAllUsers).Methods("GET")

	// Create a user
	subrouter.HandleFunc("/", controllers.CreateUser).Methods("POST")

	// serve ane image
	subrouter.HandleFunc("/image", controllers.GetImage).Methods("GET")
	// Get user by ID
	subrouter.HandleFunc("/{id}", controllers.GetUserByID).Methods("GET")

}

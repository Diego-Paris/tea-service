package routes

import (
	"net/http"

	"github.com/Diego-Paris/tea-service/pkg/controllers"
	"github.com/gorilla/mux"
)

func InitializeAllRoutes(router *mux.Router) {

	// Initialize USERS routes
	initializeUserRoutes("/users", router)

	// Catches requests with unsupported methods
	router.MethodNotAllowedHandler = http.HandlerFunc(controllers.MethodNotAllowed)

	// Catches all routes that are not declared
	router.NotFoundHandler = http.HandlerFunc(controllers.NotFound)

	// Catches all unrecovered panics that may occur
	router.Use(controllers.Recovery)
}

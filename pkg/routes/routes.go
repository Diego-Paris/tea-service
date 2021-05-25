package routes

import (
	"net/http"

	"github.com/Diego-Paris/tea-service/pkg/controllers"
	"github.com/Diego-Paris/tea-service/pkg/utils"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *http.ServeMux {

	router := mux.NewRouter()

	setRoutes("/users", UserRoutes, router)

	// Catches requests with unsupported methods
	router.MethodNotAllowedHandler = http.HandlerFunc(controllers.MethodNotAllowed)

	// Catches all routes that are not declared
	router.NotFoundHandler = http.HandlerFunc(controllers.NotFound)

	// Catches unrecovered panics that may occur
	router.Use(controllers.Recovery)

	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		req.URL.Path = utils.AddTrailingSlash(req.URL.Path)
		router.ServeHTTP(w, req)
	})
	return m
}

func setRoutes(prefix string, routes Routes, router *mux.Router) *mux.Router {
	for _, route := range routes {
		path := utils.AddTrailingSlash(prefix + route.Pattern)
		router.
			Methods(route.Method).
			Path(path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

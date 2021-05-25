package routes

import (
	"github.com/Diego-Paris/tea-service/pkg/controllers"
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

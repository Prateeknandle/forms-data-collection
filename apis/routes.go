package apis

import (
	"net/http"

	"github.com/Prateeknandle/go-plugin/services"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Post",
		"/user", // creates a user in the database
		services.PostUser,
	},
	Route{
		"GET",
		"/form", // gets all the form from the database
		services.GetForm,
	},
	Route{
		"POST",
		"/form", // creates a form in the database
		services.PostForm,
	},
	Route{
		"PUT",
		"/form", // submits a response to a form in the database
		services.PutForm,
	},
}

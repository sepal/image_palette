package web

import (
	"net/http"
	"github.com/sepal/image_palette/backend/web/handlers"
)

// Route represents a route for the web server.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes represents multiple routes.
type Routes []Route

// routes contains all routes for the application.
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"Upload",
		"POST",
		"/upload",
		handlers.Upload,
	},
	Route{
		"ImageWS",
		"GET",
		"/image/{imageID}",
		handlers.ImageWS,
	},
}

package web

import (
	"github.com/sepal/image_palette/backend/web/handlers"
	"net/http"
)

// Route represents a route for the web server.
type Route struct {
	Name        string
	Method      string
	Path        string
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
		handlers.ImageCreate,
	},
	Route{
		"Upload",
		"POST",
		"/image/{imageID}/calculated",
		handlers.ImageCalculated,
	},
}

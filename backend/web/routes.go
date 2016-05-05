package web

import (
	"github.com/sepal/image_palette/backend/web/handlers"
	"net/http"
)

type Route struct {
	Name       string
	Method     string
	Path       string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

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
}
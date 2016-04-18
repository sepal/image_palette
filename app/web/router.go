package web

import (
	"github.com/gorilla/mux"
	"github.com/sepal/color_space/app/web/handlers"
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
}

func RouteApp() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler

		handler = Logger(route.HandlerFunc, route.Name)

		r.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)

	}

	return r
}

package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

// RouteApp creates the route for the http server.
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

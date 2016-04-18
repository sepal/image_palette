package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func Route() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	return r
}

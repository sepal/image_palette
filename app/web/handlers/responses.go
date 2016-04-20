package handlers

import (
	"net/http"
	"log"
	"encoding/json"
)

func JSONResponse(w http.ResponseWriter, r *http.Request, status int,  v interface{}) {
	w.Header().Set("content-type", "application/json")

	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(v)

	if err != nil {
		log.Fatalf("Error on %v route: %v", r.URL, err)
	}
}
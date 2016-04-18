package handlers

import (
	"net/http"
	"log"
	"encoding/json"
)

func JSONResponse(w http.ResponseWriter, r *http.Request, v interface{}) {
	w.Header().Set("content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(v)

	if err != nil {
		log.Fatalf("Error on %v route: %v", r.URL, err)
	}
}
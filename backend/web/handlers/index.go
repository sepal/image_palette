package handlers

import (
	"net/http"
	"github.com/sepal/image_palette/backend/models"
)

// Index handles the / route.
func Index(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, r, http.StatusOK, models.Message{"Hello World!"});
}
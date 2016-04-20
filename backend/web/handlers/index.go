package handlers

import (
	"net/http"
	"github.com/sepal/color_space/backend/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, r, http.StatusOK, models.Message{"Hello World!"});
}
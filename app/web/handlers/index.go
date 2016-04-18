package handlers

import (
	"net/http"
	"github.com/sepal/color_space/app/web/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, r, models.Message{"Hello World!"});
}
package handlers

import (
	"net/http"
	"github.com/sepal/image_palette/backend/models"
	"log"
)

// Upload handles the route to which the images are uploaded. It accepts a POST request containing an image on the
// 'image' key and returns the meta data (id, created date...) for the uploaded image.
func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 << 20)

	file, header, err := r.FormFile("image")

	if err != nil {
		log.Printf("Error while uploading file: %v", err)
		JSONResponse(w, r, http.StatusInternalServerError, models.Message{"Failed to upload file."})
		return
	}
	defer file.Close()

	img, err := models.NewImage(file, header)

	if err != nil {
		log.Fatalf("Error while saving file %v: %v", header.Filename, err)
		JSONResponse(w, r, http.StatusInternalServerError, models.Message{"Failed to save the uploaded image."})
		return
	}

	JSONResponse(w, r, http.StatusOK, img)
}

package handlers

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/sepal/image_palette/backend/models"
	"log"
	"net/http"
)

// upgrader creates web socket for the ImageChanges function.
var upgrader = websocket.Upgrader{
	// todo: Restrict origin.
	CheckOrigin: func(r *http.Request) bool { return true },
}

// ImageCreate takes a new from the form values and returns the id of the saved image.
func ImageCreate(w http.ResponseWriter, r *http.Request) {
	// Parse the multiple part form data, allow a maximum size of 5MB.
	err := r.ParseMultipartForm(5 << 20)

	if err != nil {
		log.Printf("Could not parse form data due to: %v", err)
		JSONResponse(w, r, http.StatusBadRequest, models.Message{err.Error()})
		return
	}

	// Get the image from the key 'image'.
	file, header, err := r.FormFile("image")

	if err != nil {
		log.Printf("Error while uploading file: %v", err)
		JSONResponse(w, r, http.StatusInternalServerError, models.Message{"Failed to upload file."})
		return
	}
	defer file.Close()

	// Create a new image model, which also returns the ID.
	img, err := models.NewImage(file, header)

	if err != nil {
		log.Fatalf("Error while saving file %v: %v", header.Filename, err)
		JSONResponse(w, r, http.StatusInternalServerError, models.Message{"Failed to save the uploaded image."})
		return
	}

	JSONResponse(w, r, http.StatusOK, img)
}

// ImageChanges creates a new web socket to subscribe for an image. The server will returned a json object with the
// dominant colors, as soon as they are calculated.
func ImageCalculated(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["imageID"]
	log.Printf("Creating socket for image ID: %v", id)

	// Subscribe to changes for that create image.
	stream, err := models.ImageChanges(id)

	if err != nil {
		log.Printf("Could not subscribe to image changes with the id %v because of: %v", id, err)
		JSONResponse(w, r, http.StatusInternalServerError, models.Message{"Could not create websocket"})
		return
	}

	// Create a new web socket.
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Could not create websocket because of: %v", err)
		JSONResponse(w, r, http.StatusInternalServerError, models.Message{"Could not create websocket."})
		return
	}

	// Create a new routine to wait for calculation.
	go func() {
		defer ws.Close()

		update := make(map[string]models.Image)

		stream.Next(&update)

		// The image should have a 'new_val' containing the color scheme. That's what we want to return.
		if update["old_val"] != (models.Image{}) && update["new_val"] != (models.Image{}) {
			ws.WriteJSON(update["new_val"].Colors)
		}
	}()
}

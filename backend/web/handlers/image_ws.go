package handlers

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"github.com/sepal/image_palette/backend/models"
	"github.com/gorilla/mux"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ImageWS(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["imageID"]
	log.Printf("Creating socket for image ID: %v", id)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Could not create websocket because of: %v", err)
		JSONResponse(w, r, http.StatusInternalServerError, models.Message{"Could not create websocket"})
		return
	}

	stream, err := models.ImageChanges(id)
	if err != nil {
		log.Printf("Could not subscribe to image changes with the id %v because of: %v", id, err)
		JSONResponse(w, r, http.StatusInternalServerError, models.Message{"Could not create websocket"})
		return
	}

	go func() {
		defer ws.Close()
		update := make(map[string]models.Image)
		stream.Next(&update)
		if update["old_val"] != (models.Image{}) && update["new_val"] != (models.Image{}) {
			ws.WriteJSON(update["new_val"])
		}
	}()
}

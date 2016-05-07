package palette

import (
	"github.com/sepal/image_palette/backend/models"
	"log"
)

// Listen startes listening to incoming messages and spawns Worker functions to calculate the color schemes for the
// images.
func Listen() {
	log.Printf("Started listening")
	stream, err := models.ImageChanges()

	if err != nil {
		log.Fatalf("Error while trying to subscribe to images feed: %v", err)
	}



	go func() {
		update := make(map[string]models.Image)
		for stream.Next(&update) {
			// If old_val is empty, then we received a new Image, for which we should spawn a new worker.
			if update["old_val"] == (models.Image{}) {
				go Worker(update["new_val"])
			}
		}
	}()
}

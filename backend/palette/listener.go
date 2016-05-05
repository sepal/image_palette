package palette

import (
	r "gopkg.in/dancannon/gorethink.v2"
	"github.com/sepal/image_palette/backend/models"
	"log"
)

func Listen() {
	log.Printf("Started listening")
	stream, err := r.Table("images").Changes().Run(models.Session)

	if err != nil {
		log.Fatalf("Error while trying to subscribe to images feed: %v", err)
	}



	go func() {
		hash := make(map[string]models.Image)
		for stream.Next(&hash) {
			// If old_val is empty, then we received a new Image, for which we should spawn a new worker.
			if hash["old_val"] == (models.Image{}) {

			}
		}
	}()
}

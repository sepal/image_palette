package models

import (
	r "github.com/dancannon/gorethink"
	"github.com/satori/go.uuid"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"time"
)

var UploadDir = "/tmp"

type Image struct {
	ID       string `gorethink:"id,omitempty" json:"id"`
	Filename string `gorethink:"filename,omitempty" json:"filename"`
	Path     string `gorethink:"path,omitempty" json:"path"`
	Created  int64  `gorethink:"created,omitempty" json:"created"`
}

func NewImage(file multipart.File, header *multipart.FileHeader) (*Image, error) {
	id := uuid.NewV4().String()

	ext := path.Ext(header.Filename)

	img := &Image{
		ID:       id,
		Filename: header.Filename,
		Path:     id + ext,
	}

	log.Printf("Creating new image: %v", img.Path)

	f, err := os.OpenFile(path.Join(UploadDir, img.Path), os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = io.Copy(f, file)

	if err != nil {
		return nil, err
	}

	log.Printf("Saving image %v with ID %v to database.", img.Filename, img.ID)

	img.Created = time.Now().Unix()
	err = r.Table("images").Insert(img).Exec(session)

	return img, err
}

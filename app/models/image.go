package models

import (
	"mime/multipart"
	"log"
	"os"
	"io"
	"path"
	"github.com/satori/go.uuid"
)

var UploadDir = "/tmp"

type Image struct {
	Filename string
	Path     string
}

func NewImage(file multipart.File, header *multipart.FileHeader) (*Image, error) {
	img := Image{
		Filename: header.Filename,
		Path: path.Join(UploadDir, header.Filename),
	}

	log.Printf("Creating new image: %v", header.Header)

	f, err := os.OpenFile(img.Path, os.O_WRONLY | os.O_CREATE, 0666)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = io.Copy(f, file)

	if err != nil {
		return nil, err
	}

	return &img, nil
}
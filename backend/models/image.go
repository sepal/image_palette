package models

import (
	r "gopkg.in/dancannon/gorethink.v2"
	"github.com/satori/go.uuid"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"time"
	"net/http"
	"image"
	"image/png"
	"image/jpeg"
	"fmt"
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
	err = r.Table("images").Insert(img).Exec(Session)

	return img, err
}

func (i *Image) openFile() (*os.File, error) {
	file, err := os.Open(path.Join(UploadDir, i.Path))
	return file, err
}

func GetType(file *os.File) (string, error) {
	// Read the first 256 bytes which should be enough to detect the image type.
	bytes := make([]byte, 256)
	_, err := file.Read(bytes)

	if err != nil {
		return "", err
	}

	// Reset read pointer to 0.
	file.Seek(0, 0)

	return http.DetectContentType(bytes), nil
}

func (i *Image) GetImage() (*image.Image, error) {
	file, err := i.openFile()

	if err != nil {
		return nil, err
	}

	imgType, err := GetType(file)

	if err != nil {
		return nil, err
	}

	var img image.Image

	switch imgType {
	case "image/png":
		img, err = png.Decode(file)
		break
	case "image/jpeg":
		img, err = jpeg.Decode(file)
		break
	default:
		return nil, fmt.Errorf("Unsupported image type %v", imgType)
	}

	return &img, nil
}
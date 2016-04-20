package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/mitchellh/colorstring"
	"github.com/sepal/color_space/app/web"
	"github.com/sepal/color_space/app/models"
	"log"
	"net/http"
	"os"
)

var host, uploadDir string
var port int = 8080

// PrintError exits the program with an error.
func PrintError(err error) {
	fmt.Println(colorstring.Color("[red]" + err.Error()))
	os.Exit(1)
}

func checkArgs() {
	if port == 0 {
		port = 8080
	}

	if uploadDir != "" {
		models.UploadDir = uploadDir
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "image_colors"
	app.Usage = "A small web app, which displays a color histogram for the uploaded image."
	app.HideVersion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "interface, i",
			Usage:       "Interface on where to run the webserver",
			Destination: &host,
			EnvVar:      "INTERFACE",
		},
		cli.IntFlag{
			Name:        "port, p",
			Usage:       "Port on which the webserver should listen",
			Destination: &port,
			EnvVar:      "PORT",
		},
		cli.StringFlag{
			Name:        "uploadDir, u",
			Usage:       "Directory to store the images",
			Destination: &uploadDir,
			EnvVar:      "UPLOAD_DIR",
		},
	}

	app.Action = func(c *cli.Context) {
		checkArgs()
		str := fmt.Sprintf("%v:%v", host, port)

		log.Printf("Starting server at %v", str)
		err := http.ListenAndServe(str, web.RouteApp())

		if err != nil {
			PrintError(err)
		}
	}

	app.Run(os.Args)
}

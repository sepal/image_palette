package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/mitchellh/colorstring"
	"github.com/sepal/image_palette/backend/models"
	"github.com/sepal/image_palette/backend/palette"
	"github.com/sepal/image_palette/backend/web"
	"log"
	"net/http"
	"os"
)

var host, uploadDir, rethinkHost, database string
var port int = 8080

// PrintError exits the program with an error.
func PrintError(err error) {
	fmt.Println(colorstring.Color("[red]" + err.Error()))
	os.Exit(1)
}

// Set defaults for the main variables, which may been have set via environment variables or arguments.
func checkArgs() {
	if port == 0 {
		port = 8080
	}

	if uploadDir != "" {
		models.UploadDir = uploadDir
	}

	if rethinkHost == "" {
		rethinkHost = "localhost:28015"
	}

	if database == "" {
		database = "image_palette"
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
		cli.StringFlag{
			Name:        "rethinkHost, r",
			Usage:       "RethinkDB host to connect to",
			Destination: &rethinkHost,
			EnvVar:      "RETHINK_HOST",
		},
		cli.StringFlag{
			Name:        "database, d",
			Usage:       "RethinkDB database used store the image meta data",
			Destination: &database,
			EnvVar:      "DATABASE",
		},
	}

	app.Action = func(c *cli.Context) {
		checkArgs()

		log.Printf("Trying to connect to %v using the database %v\n", rethinkHost, database)
		err := models.Connect(rethinkHost, database)

		if err != nil {
			PrintError(err)
		}

		palette.Listen()

		host := fmt.Sprintf("%v:%v", host, port)

		log.Printf("Starting server at %v", host)
		err = http.ListenAndServe(host, web.RouteApp())

		if err != nil {
			PrintError(err)
		}
	}

	app.Run(os.Args)
}

package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/mitchellh/colorstring"
	"github.com/sepal/color_space/app/web"
	"log"
	"net/http"
	"os"
)

var host string = ""
var port int

// PrintError exits the program with an error.
func PrintError(err error) {
	fmt.Println(colorstring.Color("[red]" + err.Error()))
	os.Exit(1)
}

func checkArgs() {
	if port == 0 {
		port = 8080
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

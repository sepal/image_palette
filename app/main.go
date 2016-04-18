package main

import (
	"github.com/codegangsta/cli"
	"fmt"
	"os"
)

var host string = ""
var port int

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
		fmt.Println(str)
	}

	app.Run(os.Args)
}

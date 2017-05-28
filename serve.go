package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "serve"
	app.Version = "0.0.1"
	app.Description = "A simple HTTP server for serving static files."
	app.Authors = []cli.Author{
		{Name: "Sandeep Raju Prabhakar", Email: "me@sandeepraju.in"},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "address, a",
			Usage: "The IP address or hostname to run",
			Value: "localhost",
		},
		cli.IntFlag{
			Name:  "port, p",
			Usage: "The port to listen on",
			Value: 8888,
		},
	}
	app.Action = func(c *cli.Context) error {
		addr := c.String("address")
		port := c.Int("port")

		// TODO: check how to accept command options in urfave/cli
		dir := "./"

		http.Handle("/", http.FileServer(http.Dir(dir)))
		log.Printf("Serving %s at http://%s:%d/", dir, addr, port)
		err := http.ListenAndServe(fmt.Sprintf("%s:%d", addr, port), nil)
		if err != nil {
			log.Fatal(err)
			return err
		}

		return nil
	}

	app.Run(os.Args)
}

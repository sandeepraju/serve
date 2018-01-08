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
	app.Usage = ""
	app.Description = `A simple HTTP server for serving static files.

   See https://github.com/sandeepraju/serve for more details!

EXAMPLES:
   * Serve files in the current directory

			serve

   * Serve at port 8000 on 192.168.0.3

	    serve -a 192.168.0.3 -p 8000

	 * Serve /tmp

			serve -d /tmp`
	app.Authors = []cli.Author{
		{Name: "Sandeep Raju Prabhakar", Email: "me@sandeepraju.in"},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "directory, dir, d",
			Usage: "The directory to serve",
			Value: "./",
		},
		cli.StringFlag{
			Name:  "address, a",
			Usage: "The IP address or hostname of the interface",
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
		dir := c.String("directory")

		http.Handle("/", http.FileServer(http.Dir(dir)))
		log.Printf("Serving %s at http://%s:%d/", dir, addr, port)
		err := http.ListenAndServe(fmt.Sprintf("%s:%d", addr, port), nil)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	}

	app.Run(os.Args)
}

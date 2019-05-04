package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	flup "github.com/knakk/rdf"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Argu-cli"
	app.Version = "0.0.1"
	app.Compiled = time.Now()
	app.Usage = "Get your Argu data, straight to your favorite terminal!"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Joep Meindertsma",
			Email: "joep@argu.co",
		},
	}
	app.EnableBashCompletion = true

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "site",
			Value: "https://argu.co",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "get",
			Usage: "Looks Up the NameServers for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				resourceURL := "https://app.argu.co/u/joep.ttl"
				resp, err := http.Get(resourceURL)
				if err != nil {
					return err
				}
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				bodyString := string(bodyBytes)
				rdf := flup.Turtle
				DecodeAll(bodyString)([]Triple, error)
				if resp != nil {
					fmt.Println(bodyString)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	run(os.Args)
}

func run(args []string) {
	app := cli.NewApp()
	app.Name = "ldt"
	app.Version = "0.0.2"
	app.Compiled = time.Now()
	app.Usage = "Get your RDF data, straight to your favorite terminal!"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Joep Meindertsma",
			Email: "joep@argu.co",
		},
	}
	app.EnableBashCompletion = true

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "resource",
			Usage: "The URL of the resource to be fetched. Should return an N-Quads file",
		},
		cli.StringFlag{
			Name:  "subject",
			Usage: "The URL of the subject to be matched",
		},
		cli.StringFlag{
			Name:  "predicate",
			Usage: "The URL of the predicate to be matched",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "get",
			Usage: "Fetch an RDF resource",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				// resourceURL := c.Args()
				resourceURL := c.String("resource")
				subject := c.String("subject")
				object := c.String("object")
				predicate := c.String("predicate")
				resp, err := http.Get(resourceURL)
				if err != nil {
					return err
				}
				// escapedSubject := regexp.QuoteMeta(subject)
				// subjectFinder := fmt.Sprintf("(<%s> <.+?>) ([^<]+)", escapedSubject)
				// mySubject, err := regexp.Compile(subjectFinder)
				allTriples, err := Parse(resp.Body)
				// hits := findByPredicate(allTriples, predicate)
				hits := filterTriples(allTriples, subject, predicate, object)
				if len(hits) == 0 {
					log.Fatal("Not found")
				} else if hits[0] == nil {
					log.Fatal("Found, but no object in triple")
				} else {
					fmt.Println(hits[0].object)
				}
				return nil
			},
		},
	}

	err := app.Run(args)
	if err != nil {
		log.Fatal(err)
	}
}

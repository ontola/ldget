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
	app.Name = "Linked Data Get"
	app.Version = "0.0.2"
	app.Compiled = time.Now()
	app.Usage = "Get your RDF data, straight to your favorite terminal! Flags have precedence over arguments."
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Joep Meindertsma",
			Email: "joep@argu.co",
		},
	}
	app.EnableBashCompletion = true

	initialize()

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "resource, r",
			Usage: "The URL of the resource to be fetched. URL should return an N-Quads file. If this is empty, the Subject is used.",
		},
		cli.StringFlag{
			Name:  "subject, s",
			Usage: "The IRI of the subject to be matched",
		},
		cli.StringFlag{
			Name:  "predicate, p",
			Usage: "Filter by predicate",
		},
		cli.StringFlag{
			Name:  "object, o",
			Usage: "Filter by object value",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "objects",
			Aliases: []string{"o"},
			Usage:   "Fetch an RDF resource, return the values. First argument is Subject, second is Predicate.",
			Flags:   myFlags,
			Action: func(c *cli.Context) error {
				args := getArgs(c)
				resp, err := http.Get(args.resourceURL)
				if err != nil {
					return err
				}
				allTriples, err := Parse(resp.Body)
				hits := filterTriples(allTriples, args.subject, args.predicate, args.object)
				if len(hits) == 0 {
					log.Fatal("Not found")
				} else if hits[0] == nil {
					log.Fatal("Found, but no object in triple")
				} else {
					for _, element := range hits {
						fmt.Println(element.object)
					}
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

type args struct {
	resourceURL string
	subject     string
	object      string
	predicate   string
}

func getArgs(c *cli.Context) args {
	var arguments args

	arguments.subject = c.Args().Get(0)
	arguments.predicate = c.Args().Get(1)
	arguments.object = c.Args().Get(2)

	if c.String("subject") != "" {
		arguments.subject = c.String("subject")
	}
	if c.String("predicate") != "" {
		arguments.predicate = c.String("predicate")
	}
	if c.String("object") != "" {
		arguments.object = c.String("object")
	}
	arguments.subject = Mapper(arguments.subject)
	if c.String("resource") != "" {
		arguments.resourceURL = c.String("resource")
	} else {
		if arguments.subject == "" {
			log.Fatal("No resource or subject provided. See --help.")
		}
		// TODO: use content negotiation
		arguments.resourceURL = fmt.Sprintf("%v.nq", arguments.subject)
	}
	arguments.predicate = Mapper(arguments.predicate)

	return arguments
}

func initialize() {
	getAllMaps()
}

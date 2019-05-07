package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
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

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "resource, r",
			Usage: "The URL of the resource to be fetched. URL should return an N-Quads file. If this is empty, the Subject is used.",
		},
		cli.StringFlag{
			Name:  "subject, s",
			Usage: "Filter by subject IRI.  Prefixes allowed.",
		},
		cli.StringFlag{
			Name:  "predicate, p",
			Usage: "Filter by predicate IRI. Prefixes allowed.",
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
			Usage:   "Fetch an RDF resource, return the object values. First argument filters by Subject, second by Predicate.",
			Flags:   myFlags,
			Action: func(c *cli.Context) error {
				args := getArgs(c)
				resp, format, err := Negotiator(args.resourceURL)
				if err != nil {
					return err
				}
				allTriples, err := Parse(resp.Body, format)
				if err != nil {
					return err
				}
				hits := filterTriples(allTriples, args.subject, args.predicate, args.object)
				if len(hits) == 0 {
					log.Fatal("No triple found that matches your query")
				} else {
					for _, element := range hits {
						fmt.Println(element.Obj)
					}
				}
				return nil
			},
		},

		{
			Name:  "prefixes",
			Usage: "Shows your user defined prefixes from  `~/.ldget/prefixes`.",
			Action: func(c *cli.Context) error {
				for _, mapItem := range getAllMaps() {
					w := new(tabwriter.Writer)
					w.Init(os.Stdout, 15, 8, 0, '\t', 0)
					fmt.Fprintf(w, "%v\t%v\t\n", mapItem.key, mapItem.url)
					w.Flush()
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
		arguments.resourceURL = arguments.subject
	}
	arguments.predicate = Mapper(arguments.predicate)

	return arguments
}

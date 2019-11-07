package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/knakk/rdf"
	"github.com/urfave/cli"
)

// Overwrite these using ldflags
var version = fmt.Sprintf("dev%v", time.Now().Format(time.RFC3339))

func main() {
	run(os.Args)
}

func run(args []string) {
	app := cli.NewApp()
	app.Name = "ldget"
	app.Version = version
	app.Compiled = time.Now()
	app.Usage = "Get your RDF data, straight to your favorite terminal! Flags have precedence over arguments."
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Joep Meindertsma",
			Email: "joep@ontola.io",
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
			Usage: "Filter by object value.",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "triples",
			Aliases: []string{"t", ""},
			Usage:   "Fetch an RDF resource, return the triples. `ldget t ?s ?p ?o`",
			UsageText: "Fetch an RDF resource, return the triples. `ldget t ?s ?p ?o` \n" +
				"You can use . as a wildcard. \n" +
				"e.g. `ldget t dbpedia:Utrecht . dbpedia:Netherlands`",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				args := getArgs(c)
				hits := getTriples(args)
				encoder := rdf.NewTripleEncoder(os.Stdout, rdf.NTriples)
				encoder.GenerateNamespaces = true
				encoder.EncodeAll(hits)
				encoder.Close()
				return nil
			},
		},
		{
			Name:    "predicates",
			Aliases: []string{"p", ""},
			Usage:   "Fetch an RDF resource, return the predicates. `ldget p ?s ?p ?o`",
			Flags:   myFlags,
			Action: func(c *cli.Context) error {
				args := getArgs(c)
				hits := getTriples(args)
				for _, element := range hits {
					fmt.Println(element.Pred.Serialize(rdf.NTriples))
				}
				return nil
			},
		},
		{
			Name:    "objects",
			Aliases: []string{"o"},
			Usage:   "Fetch an RDF resource, return the object values. `ldget o ?s ?p ?o`",
			Flags:   myFlags,
			Action: func(c *cli.Context) error {
				args := getArgs(c)
				hits := getTriples(args)
				for _, element := range hits {
					fmt.Println(element.Obj.Serialize(rdf.NTriples))
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
		{
			Name:    "expand",
			Aliases: []string{"x"},
			Usage:   "Expands any prefix. `ldget x schema` => https://schema.org/",
			Action: func(c *cli.Context) error {
				prefix := c.Args().Get(0)
				match := prefixMap(prefix)
				if match == prefix {
					fmt.Printf("Prefix '%v' Not found \n", match)
				} else {
					fmt.Printf("%v\n", match)
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

// Check if the input string should be interpreted as a wildcard
func cleanUpArg(s string) string {
	blankArgs := map[string]bool{
		"": true,
		// Does not work in zsh, formula?
		// "*":    true,
		// Does not work in zsh, wildcard
		// "?":    true,
		".":    true,
		"null": true,
		"nil":  true,
	}

	if blankArgs[s] {
		return ""
	}

	return s
}

func getArgs(c *cli.Context) args {
	var arguments args

	subject := c.Args().Get(0)
	predicate := c.Args().Get(1)
	object := c.Args().Get(2)

	if c.String("subject") != "" {
		subject = c.String("subject")
	}
	if c.String("predicate") != "" {
		predicate = c.String("predicate")
	}
	if c.String("object") != "" {
		object = c.String("object")
	}

	arguments.subject = cleanUpArg(subject)
	arguments.predicate = cleanUpArg(predicate)
	arguments.object = cleanUpArg(object)

	arguments.subject = prefixMap(arguments.subject)
	if c.String("resource") != "" {
		arguments.resourceURL = c.String("resource")
	} else {
		if arguments.subject == "" {
			log.Fatal("No resource or subject provided. See --help.")
		}
		arguments.resourceURL = arguments.subject
	}
	arguments.predicate = prefixMap(arguments.predicate)

	return arguments
}

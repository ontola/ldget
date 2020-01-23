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

var defaultPrefixPath = `~/.ldget/prefixes`

func main() {
	run(os.Args)
}

func run(args []string) {
	app := cli.NewApp()
	app.Name = "ldget"
	app.Version = version
	app.Compiled = time.Now()
	app.Usage = "Get your RDF data, straight to your favorite terminal! Filter triples using `?s ?p ?o`."
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
		cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "Turn on verbose output, including requests and responses.",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "triples",
			Aliases: []string{"t"},
			Usage:   "Fetch an RDF resource, return the triples. Serialized as N-Triples.",
			UsageText: "`ldget t ?s ?p ?o` \n" +
				"   You can use . as a wildcard. \n" +
				"   e.g. `ldget t dbpedia:Utrecht . dbpedia:Netherlands`",
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
			Aliases: []string{"p"},
			Usage:   "Fetch an RDF resource, return the predicates.",
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
			Usage:   "Fetch an RDF resource, return the objects.",
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
			Name:    "subjects",
			Aliases: []string{"s"},
			Usage:   "Fetch an RDF resource, return the subjects.",
			Flags:   myFlags,
			Action: func(c *cli.Context) error {
				args := getArgs(c)
				hits := getTriples(args)
				for _, element := range hits {
					fmt.Println(element.Subj.Serialize(rdf.NTriples))
				}
				return nil
			},
		},
		{
			Name:    "predicateObjects",
			Aliases: []string{"po"},
			Usage:   "Fetch an RDF resource, return the predicate and object values.",
			Flags:   myFlags,
			Action: func(c *cli.Context) error {
				args := getArgs(c)
				hits := getTriples(args)
				for _, element := range hits {
					fmt.Printf("%v %v\n", element.Pred.Serialize(rdf.NTriples), element.Obj.Serialize(rdf.NTriples))
				}
				return nil
			},
		},
		{
			Name:    "show",
			Aliases: []string{"sh"},
			Usage:   "Fetch an RDF resource, return the predicate and object values in a nice table.",
			Flags:   myFlags,
			Action: func(c *cli.Context) error {
				args := getArgs(c)
				hits := getTriples(args)
				w := new(tabwriter.Writer)
				w.Init(os.Stdout, 0, 8, 2, '\t', 0)
				for _, element := range hits {
					// fmt.Printf("%v %v\n", element.Pred.Serialize(rdf.NTriples), element.Obj.Serialize(rdf.NTriples))
					fmt.Fprintf(w, "%v\t%v\t\n", tryURLToPrefix(element.Pred.Serialize(rdf.NTriples), args), element.Obj.Serialize(rdf.NTriples))
				}
				w.Flush()
				return nil
			},
		},
		{
			Name:  "prefixes",
			Usage: fmt.Sprintf("Shows your user defined prefixes from  '%v'.", defaultPrefixPath),
			Action: func(c *cli.Context) error {
				for _, mapItem := range getAllPrefixes() {
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
				args := getArgs(c)
				match := prefixToURL(prefix, args.prefixes)
				if match == prefix {
					fmt.Printf("Prefix '%v' not found in '%v'  \n", match, defaultPrefixPath)
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

// Args - All instance specific arguments
type Args struct {
	resourceURL string
	subject     string
	object      string
	predicate   string
	verbose     bool
	prefixes    []Prefix
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

func getArgs(c *cli.Context) Args {
	var arguments Args

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

	arguments.prefixes = getAllPrefixes()

	arguments.subject = prefixToURL(arguments.subject, arguments.prefixes)
	if c.String("resource") != "" {
		arguments.resourceURL = c.String("resource")
	} else {
		if arguments.subject == "" {
			log.Fatal("No resource or subject provided. See --help.")
		}
		arguments.resourceURL = arguments.subject
	}
	arguments.predicate = prefixToURL(arguments.predicate, arguments.prefixes)
	if c.Bool("verbose") == true {
		arguments.verbose = true
	}
	return arguments
}

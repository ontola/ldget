package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/urfave/cli"
)

func main() {
	run(os.Args)
}

func run(args []string) {
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
				resourceURL := "https://app.argu.co/u/joep.nq"
				subject := "https://app.argu.co/argu/u/joep"
				// predicate := "http://schema.org/name"
				resp, err := http.Get(resourceURL)
				if err != nil {
					return err
				}

				scanner := bufio.NewScanner(bufio.NewReader(resp.Body))
				// Iterates over every single triple
				for scanner.Scan() {
					triple := scanner.Text()
					escapedSubject := regexp.QuoteMeta(subject)
					subjectFinder := fmt.Sprintf("(<%s> <.+?>) ([^<]+)", escapedSubject)
					mySubject, err := regexp.Compile(subjectFinder)
					if mySubject.MatchString(triple) {
						fmt.Println("Your sub")
						findObject, err := regexp.Compile(`(<.+?> <.+?>) ([^<]+)`)
						if err != nil {
							return err
						}
						matches := findObject.FindStringSubmatch(triple)
						fmt.Println(matches[2], triple)
					}
					if err != nil {
						return err
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

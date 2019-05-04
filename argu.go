package main

import (
	"fmt"
	// "io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	rdf "github.com/knakk/rdf"
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
				// resourceURL := "https://app.argu.co/u/joep.n3dawd"
				resourceURL := "https://gist.githubusercontent.com/kal/ee1260ceb462d8e0d5bb/raw/1364c2bb469af53323fdda508a6a579ea60af6e4/log_sample.ttl"
				resp, err := http.Get(resourceURL)
				if err != nil {
					return err
				}
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				bodyString := string(bodyBytes)
				if len(bodyString) > 0 {
					fmt.Println("Body present")
				}
				dec := rdf.NewTripleDecoder(resp.Body, rdf.NTriples)
				fmt.Println("Start decoding...")
				triples, err := dec.DecodeAll()
				fmt.Println(triples)
				fmt.Println(err)
				// for triple, err := dec.Decode(); err != io.EOF; triple, err = dec.Decode() {
				// 	fmt.Println("In the loop!")
				// 	fmt.Println(err)
				// 	fmt.Println(triple)
				// }
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

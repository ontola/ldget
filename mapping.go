package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"regexp"
)

type prefix struct {
	key string
	url string
}

var selector, _ = regexp.Compile(`(.*)=(.*)`)

func readMap(filePath string) []prefix {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var prefixes []prefix
	scanner := bufio.NewScanner(file) // f is the *os.File
	for scanner.Scan() {
		line := scanner.Text()
		matches := selector.FindStringSubmatch(line)
		var p prefix
		if len(matches) < 2 {
			log.Fatal("Something is wrong with your mapping file.")
		}
		p.key = matches[1]
		p.url = matches[2]
		prefixes = append(prefixes, p)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return prefixes
}

func getAllMaps() []prefix {
	var allPrefixes []prefix
	allPrefixes = append(allPrefixes, readMap("defaultMapping")...)

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	userMappingLocation := fmt.Sprintf("%v/.ldget/mapping", usr.HomeDir)

	allPrefixes = append(allPrefixes, readMap(userMappingLocation)...)
	return allPrefixes
}

// Mapper -- converts a mappeed string to a URI
func Mapper(str string) string {
	matched, err := regexp.MatchString(`http.*`, str)
	if err != nil {
		log.Fatal(err)
	}
	output := str
	if matched {
		return output
	}
	for _, prefix := range getAllMaps() {
		if prefix.key == str {
			output = prefix.url
			break
		}
	}
	return output
}

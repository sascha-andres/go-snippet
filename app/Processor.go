package app

import (
	"log"
	"regexp"

	"github.com/sascha-andres/go-snippet/data"
	"github.com/spf13/viper"
)

var (
	snippetRegexp *regexp.Regexp
)

// Processor parses a file for marks and replaces them with the snippet value
func Processor(snippetMarker data.SnippetMarker) error {
	if viper.GetBool("verbose") {
		log.Printf("Snippet directory: '%s'", viper.GetString("snippetdir"))
		log.Printf("File: '%s'", viper.GetString("file"))
	}
	checkExistence(viper.GetString("snippetdir"))
	checkExistence(viper.GetString("file"))
	snippetRegexp = regexp.MustCompile(string(snippetMarker))
	var (
		content []string
		err     error
		marker  []data.SnippetMarkerLocation
	)
	if content, err = readInputFile(); err != nil {
		return err
	}
	if marker, err = findMarker(content); err != nil {
		return err
	}
	if viper.GetBool("verbose") {
		for i, mark := range marker {
			log.Printf("%d: %#v\n", i, mark)
		}
	}
	return saveFile(content, marker)
}

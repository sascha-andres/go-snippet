package app

import (
	"log"
	"strings"

	"github.com/sascha-andres/go-snippet/data"
	"github.com/spf13/viper"
)

func extractSnippetFromMarker(marker data.SnippetMarker) (string, error) {
	if viper.GetBool("verbose") {
		log.Printf("extractSnippetFromMarker: '%s'\n", string(marker))
	}
	return strings.Trim(snippetRegexp.FindStringSubmatch(string(marker))[1], " "), nil
}

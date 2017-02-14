package app

import (
	"log"

	"github.com/sascha-andres/go-snippet/data"
	"github.com/spf13/viper"
)

func findMarker(lines []string) ([]data.SnippetMarkerLocation, error) {
	if viper.GetBool("verbose") {
		log.Println("findMarker")
	}
	var result []data.SnippetMarkerLocation
	for lineNumber, line := range lines {
		if snippetRegexp.MatchString(line) {
			found := false
			for index, mark := range result {
				if string(mark.Marker) == line && mark.StopLine == 0 {
					found = true
					mark.StopLine = lineNumber
					result[index] = mark
					break
				}
			}
			if !found {
				result = append(result, data.SnippetMarkerLocation{Marker: data.SnippetMarker(line), StartLine: lineNumber})
			}
		}
	}
	return result, nil
}

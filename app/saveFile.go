package app

import (
	"log"

	"github.com/sascha-andres/go-templating/data"
	"github.com/spf13/viper"
)

func saveFile(lines []string, marker []data.SnippetMarkerLocation) error {
	if viper.GetBool("verbose") {
		log.Println("saveFile")
	}
	var (
		writePointer int
		result       []string
	)
	writePointer = 0
	for _, mark := range marker {
		for writePointer <= mark.StartLine {
			result = append(result, lines[writePointer])
			writePointer++
		}
		snippet, err := extractSnippetFromMarker(mark.Marker)
		if err != nil {
			log.Fatal(err)
		}
		snippetContent, err := getSnippetContent(snippet)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, snippetContent)
		result = append(result, string(mark.Marker))
		if mark.StopLine > 0 {
			writePointer = mark.StopLine + 1
		}
	}
	for writePointer < len(lines) {
		result = append(result, lines[writePointer])
		writePointer++
	}
	err := writeLines(result, viper.GetString("file"))
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

package app

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/viper"
)

func getSnippetContent(snippet string) (string, error) {
	if viper.GetBool("verbose") {
		log.Printf("getSnippetContent: '%s'\n", snippet)
	}
	path := fmt.Sprintf("%s/%s", viper.GetString("templatedir"), snippet)
	if ok, err := exists(path); !ok || err != nil {
		return "", fmt.Errorf("'%s' does not exist", snippet)
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

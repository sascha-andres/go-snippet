package app

import (
	"log"

	"github.com/spf13/viper"
)

func checkExistence(path string) {
	if viper.GetBool("verbose") {
		log.Printf("checkExistence: '%s'\n", path)
	}
	if exists, err := exists(path); err != nil || !exists {
		if !exists {
			log.Fatalf("%s does not exist", path)
		} else {
			log.Fatal(err)
		}
	}
}

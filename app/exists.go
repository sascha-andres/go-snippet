package app

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func exists(path string) (bool, error) {
	if viper.GetBool("verbose") {
		log.Printf("exists: '%s'\n", path)
	}
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

package app

import (
	"bufio"
	"log"
	"os"

	"github.com/spf13/viper"
)

func readInputFile() ([]string, error) {
	if viper.GetBool("verbose") {
		log.Println("readInputFile")
	}
	file, err := os.OpenFile(viper.GetString("file"), os.O_RDONLY, 0400)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if viper.GetBool("verbose") {
		log.Printf("Read %d lines", len(lines))
	}
	return lines, nil
}

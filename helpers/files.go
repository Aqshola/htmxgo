package helpers

import (
	"log"
	"os"
)

func GetTemplatePath(path string) string {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return wd + path
}

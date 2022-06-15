package main

import (
	"log"
	"os"

	"github.com/asraa/slsa-example/parse"
)

func main() {
	file, err := os.ReadFile("test.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err := parse.ParseYaml(file); err != nil {
		log.Fatal(err)
	}
}

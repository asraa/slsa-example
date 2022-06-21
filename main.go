package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/asraa/slsa-example/parse"
)

var input = flag.String("input", "test.yaml", "input YAML file to parse")

// This is a test program that parses an input YAML file.
func main() {
	flag.Parse()

	fmt.Printf("::set-output name=egg::run-me-at-the-attested-branch\n")

	file, err := os.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}
	if err := parse.ParseYaml(file); err != nil {
		log.Fatal(err)
	}
}

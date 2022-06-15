package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func ParseYaml(b []byte) error {
	var t interface{}
	return yaml.Unmarshal(b, &t)

}

func main() {
	file, err := os.ReadFile("test.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err := ParseYaml(file); err != nil {
		log.Fatal(err)
	}
}

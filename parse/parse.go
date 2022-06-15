package parse

import "gopkg.in/yaml.v3"

func ParseYaml(b []byte) error {
	var t interface{}
	return yaml.Unmarshal(b, &t)
}

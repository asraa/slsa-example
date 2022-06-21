package parse

import (
	"github.com/google/martian/v3/log"
	"gopkg.in/yaml.v3"
)

func ParseYaml(b []byte) error {
	var t interface{}
	if err := yaml.Unmarshal(b, &t); err != nil {
		log.Errorf("error parsing\n:%s", b)
		return err
	}
	return nil
}

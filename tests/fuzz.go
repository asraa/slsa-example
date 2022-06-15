//go:build gofuzz
// +build gofuzz

package tests

import (
	"testing"

	"github.com/asraa/slsa-example/parse"
)

func FuzzParseYaml(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		if err := parse.ParseYaml(data); err != nil {
			t.Errorf("%s %v", data, err)
		}
	})
}

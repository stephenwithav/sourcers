package yaml

import (
	"io"

	"github.com/ardanlabs/conf"
	"github.com/stephenwithav/sourcers"
	"gopkg.in/yaml.v3"
)

// NewSource returns a conf.Sourcer and, potentially, an error if a
// read error occurs or the Reader contains an invalid YAML document.
func NewSource(r io.Reader) (conf.Sourcer, error) {
	return sourcers.From(yaml.Unmarshal, r)
}

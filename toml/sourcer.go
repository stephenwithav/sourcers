package toml

import (
	"io"

	"github.com/BurntSushi/toml"
	"github.com/ardanlabs/conf"
	"github.com/stephenwithav/sourcers"
)

// NewSource returns a conf.Sourcer and, potentially, an error if a
// read error occurs or the Reader contains an invalid TOML document.
func NewSource(r io.Reader) (conf.Sourcer, error) {
	return sourcers.From(toml.Unmarshal, r)
}

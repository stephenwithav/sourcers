package json

import (
	"encoding/json"
	"io"

	"github.com/ardanlabs/conf"
	"github.com/stephenwithav/sourcers"
)

// NewSource returns a conf.Sourcer and, potentially, an error if a
// read error occurs or the Reader contains an invalid JSON document.
func NewSource(r io.Reader) (conf.Sourcer, error) {
	return sourcers.From(json.Unmarshal, r)
}

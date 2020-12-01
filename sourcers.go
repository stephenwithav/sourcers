package sourcers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/ardanlabs/conf"
	"gopkg.in/yaml.v3"
)

type genericSourcer struct {
	m map[string]string
}

func (s *genericSourcer) Source(fld conf.Field) (string, bool) {
	if fld.Options.ShortFlagChar != 0 {
		flagKey := fld.Options.ShortFlagChar
		k := strings.ToLower(string(flagKey))
		if val, found := s.m[k]; found {
			return val, found
		}
	}

	k := strings.ToLower(strings.Join(fld.FlagKey, `_`))
	val, found := s.m[k]
	return val, found
}

type unmarshalFunc func([]byte, interface{}) error

// sourceFrom ...
func sourceFrom(f unmarshalFunc, r io.Reader) (conf.Sourcer, error) {
	if r == nil {
		return &genericSourcer{m: nil}, nil
	}

	src, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	tmpMap := make(map[string]interface{})
	err = f(src, &tmpMap)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	for key, value := range tmpMap {
		switch v := value.(type) {
		case float64:
			m[key] = strings.TrimRight(fmt.Sprintf("%f", v), "0.")
		case bool:
			m[key] = fmt.Sprintf("%t", v)
		case string:
			m[key] = value.(string)
		}
	}

	return &genericSourcer{m: m}, nil
}

// NewSourceJSON accepts a reader containing a valid JSON struct and
// returns a ardanlabs/conf.Sourcer for use by that package's Parse
// method.
func NewSourceJSON(r io.Reader) (conf.Sourcer, error) {
	return sourceFrom(json.Unmarshal, r)
}

// NewSourceYAML accepts a reader containing a valid YAML struct and
// returns a ardanlabs/conf.Sourcer for use by that package's Parse
// method.
func NewSourceYAML(r io.Reader) (conf.Sourcer, error) {
	return sourceFrom(yaml.Unmarshal, r)
}

// NewSourceTOML accepts a reader containing a valid TOML struct and
// returns a ardanlabs/conf.Sourcer for use by that package's Parse
// method.
func NewSourceTOML(r io.Reader) (conf.Sourcer, error) {
	return sourceFrom(toml.Unmarshal, r)
}

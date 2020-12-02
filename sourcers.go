package sourcers

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/ardanlabs/conf"
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

type UnmarshalFunc func([]byte, interface{}) error

// From returns a conf.Sourcer and an error when receiving an
// UnmarshalFunc and an io.Reader containing the configuration in the
// requested format.
//
// This is a helper func for subpackages and custom formats.
func From(f UnmarshalFunc, r io.Reader) (conf.Sourcer, error) {
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

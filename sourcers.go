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

type JSON struct {
	m map[string]string
}

// NewSourceJSON accepts a reader containing a valid JSON struct and
// returns a ardanlabs/conf.Sourcer for use by that package's Parse
// method.
func NewSourceJSON(r io.Reader) (*JSON, error) {
	src, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	tmpMap := make(map[string]interface{})
	err = json.Unmarshal(src, &tmpMap)
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

	return &JSON{m: m}, nil
}

// Source implements the conf.Sourcer interface. It returns the
// stringified value stored at the specified key in the JSON>
func (j *JSON) Source(fld conf.Field) (string, bool) {
	if fld.Options.ShortFlagChar != 0 {
		flagKey := fld.Options.ShortFlagChar
		k := strings.ToLower(string(flagKey))
		if val, found := j.m[k]; found {
			return val, found
		}
	}

	k := strings.ToLower(strings.Join(fld.FlagKey, `_`))
	val, found := j.m[k]
	return val, found
}

type YAML struct {
	m map[string]string
}

// NewSourceYAML accepts a reader containing a valid YAML struct and
// returns a ardanlabs/conf.Sourcer for use by that package's Parse
// method.
func NewSourceYAML(r io.Reader) (*YAML, error) {
	src, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	tmpMap := make(map[string]interface{})
	err = yaml.Unmarshal(src, &tmpMap)
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

	return &YAML{m: m}, nil
}

// Source implements the conf.Sourcer interface. It returns the
// stringified value stored at the specified key in the YAML>
func (j *YAML) Source(fld conf.Field) (string, bool) {
	if fld.Options.ShortFlagChar != 0 {
		flagKey := fld.Options.ShortFlagChar
		k := strings.ToLower(string(flagKey))
		if val, found := j.m[k]; found {
			return val, found
		}
	}

	k := strings.ToLower(strings.Join(fld.FlagKey, `_`))
	val, found := j.m[k]
	return val, found
}

type TOML struct {
	m map[string]string
}

// NewSourceTOML accepts a reader containing a valid TOML struct and
// returns a ardanlabs/conf.Sourcer for use by that package's Parse
// method.
func NewSourceTOML(r io.Reader) (*TOML, error) {
	src, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	tmpMap := make(map[string]interface{})
	err = toml.Unmarshal(src, &tmpMap)
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

	return &TOML{m: m}, nil
}

// Source implements the conf.Sourcer interface. It returns the
// stringified value stored at the specified key in the TOML>
func (j *TOML) Source(fld conf.Field) (string, bool) {
	if fld.Options.ShortFlagChar != 0 {
		flagKey := fld.Options.ShortFlagChar
		k := strings.ToLower(string(flagKey))
		if val, found := j.m[k]; found {
			return val, found
		}
	}

	k := strings.ToLower(strings.Join(fld.FlagKey, `_`))
	val, found := j.m[k]
	return val, found
}

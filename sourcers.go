package sourcers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/ardanlabs/conf"
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

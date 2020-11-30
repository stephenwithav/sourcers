# sourcers

Additional `Sourcers` for [ardanlabs/conf](https://github.com/ardanlabs/conf/).

## Interface definition

```go
// Parse parses configuration into the provided struct.
func Parse(args []string, namespace string, cfgStruct interface{}, sources ...Sourcer) error {

// Sourcer provides the ability to source data from a configuration source.
// Consider the use of lazy-loading for sourcing large datasets or systems.
type Sourcer interface {

	// Source takes the field key and attempts to locate that key in its
	// configuration data. Returns true if found with the value.
	Source(fld Field) (string, bool)
}
```

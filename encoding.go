package encoding

import (
	"github.com/go-zoox/encoding/json"
	"github.com/go-zoox/encoding/toml"
	"github.com/go-zoox/encoding/yaml"
)

type Encoding interface {
	// Encode returns the X encoding of the data.
	Encode([]byte) []byte
	// Decode decodes the X-encoded data and stores the result in the value pointed to by v.
	Decode([]byte) ([]byte, error)
}

var Json json.JSON
var Yaml yaml.YAML
var Toml toml.TOML

package encoding

import (
	"github.com/go-zoox/encoding/json"
	"github.com/go-zoox/encoding/toml"
	"github.com/go-zoox/encoding/yaml"
)

// Encoding is an interface for encoding/decoding.
type Encoding interface {
	// Encode returns the X encoding of the data.
	Encode([]byte) []byte
	// Decode decodes the X-encoded data and stores the result in the value pointed to by v.
	Decode([]byte) ([]byte, error)
}

// Json exports from encoding/json
var Json json.JSON

// Yaml exports from encoding/yaml
var Yaml yaml.YAML

// Toml exports from encoding/toml
var Toml toml.TOML

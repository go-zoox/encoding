package yaml

import (
	goyaml "github.com/goccy/go-yaml"
)

type YAML struct {
}

func New() *YAML {
	return &YAML{}
}

// Encode returns the YAML encoding of the data.
func (j *YAML) Encode(v interface{}) ([]byte, error) {
	return goyaml.Marshal(v)
}

// Decode decodes the YAML-encoded data and stores the result in the value pointed to by v.
func (j *YAML) Decode(raw []byte, v interface{}) error {
	return goyaml.Unmarshal(raw, v)
}

// Encode returns the YAML encoding of the data.
//	Static Method
func Encode(v interface{}) ([]byte, error) {
	return New().Encode(v)
}

// Decode decodes the YAML-encoded data and stores the result in the value pointed to by v.
//	Static Method
func Decode(raw []byte, v interface{}) error {
	return New().Decode(raw, v)
}

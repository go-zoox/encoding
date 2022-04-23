package json

import (
	goini "github.com/go-zoox/ini"
)

type INI struct {
}

func New() *INI {
	return &INI{}
}

// Encode returns the JSON encoding of the data.
func (j *INI) Encode(v interface{}) ([]byte, error) {
	return goini.Marshal(v)
}

// Decode decodes the JSON-encoded data and stores the result in the value pointed to by v.
func (j *INI) Decode(raw []byte, v interface{}) error {
	return goini.Unmarshal(raw, v)
}

// Encode returns the JSON encoding of the data.
//	Static Method
func Encode(v interface{}) ([]byte, error) {
	return New().Encode(v)
}

// Decode decodes the JSON-encoded data and stores the result in the value pointed to by v.
//	Static Method
func Decode(raw []byte, v interface{}) error {
	return New().Decode(raw, v)
}

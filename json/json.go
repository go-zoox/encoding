package json

import (
	gojson "encoding/json"
)

type JSON struct {
}

func New() *JSON {
	return &JSON{}
}

// Encode returns the JSON encoding of the data.
func (j *JSON) Encode(v interface{}) ([]byte, error) {
	return gojson.Marshal(v)
}

// Decode decodes the JSON-encoded data and stores the result in the value pointed to by v.
func (j *JSON) Decode(raw []byte, v interface{}) error {
	return gojson.Unmarshal(raw, v)
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

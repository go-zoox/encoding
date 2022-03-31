package toml

import (
	goini "github.com/pelletier/go-toml"
)

type TOML struct {
}

func New() *TOML {
	return &TOML{}
}

// Encode returns the TOML encoding of the data.
func (j *TOML) Encode(v interface{}) ([]byte, error) {
	return goini.Marshal(v)
}

// Decode decodes the TOML-encoded data and stores the result in the value pointed to by v.
func (j *TOML) Decode(raw []byte, v interface{}) error {
	return goini.Unmarshal(raw, v)
}

// Encode returns the TOML encoding of the data.
//	Static Method
func Encode(v interface{}) ([]byte, error) {
	return New().Encode(v)
}

// Decode decodes the TOML-encoded data and stores the result in the value pointed to by v.
//	Static Method
func Decode(raw []byte, v interface{}) error {
	return New().Decode(raw, v)
}

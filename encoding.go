package encoding

// Encoding is an interface for encoding/decoding.
type Encoding interface {
	// Encode returns the X encoding of the data.
	Encode(v interface{}) ([]byte, error)
	// Decode decodes the X-encoded data and stores the result in the value pointed to by v.
	Decode(raw []byte, v interface{})
}

package encoding

// Encoding is an interface for encoding/decoding.
type Encoding interface {
	// Encode returns the X encoding of the data.
	Encode([]byte) []byte
	// Decode decodes the X-encoded data and stores the result in the value pointed to by v.
	Decode([]byte) ([]byte, error)
}

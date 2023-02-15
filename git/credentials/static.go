package credentials

// Encode returns the git-credentials encoding of the data.
//
//	Static Method
func Encode(v map[string]*Item) ([]byte, error) {
	return New().Encode(v)
}

// Decode decodes the git-credentials-encoded data and stores the result in the value pointed to by v.
//
//	Static Method
func Decode(raw []byte) (map[string]*Item, error) {
	return New().Decode(raw)
}

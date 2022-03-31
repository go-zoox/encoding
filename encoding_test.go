package encoding

import (
	"reflect"
	"testing"
)

func TestEncoding(t *testing.T) {
	obj := map[string]interface{}{
		"foo": "bar",
		"baz": "qux",
	}

	encoded, err := Json.Encode(obj)
	if err != nil {
		t.Errorf("error encoding: %s", err)
	}

	decoded := map[string]interface{}{}
	err = Json.Decode(encoded, &decoded)
	if err != nil {
		t.Errorf("error decoding: %s", err)
	}

	if !reflect.DeepEqual(obj, decoded) {
		t.Errorf("expected %v, got %v", obj, decoded)
	}

	// fmt.Println("obj:", obj)
	// fmt.Println("decoded:", decoded)
}

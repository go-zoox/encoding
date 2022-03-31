package json

import (
	"reflect"
	"testing"
)

func TestJSON(t *testing.T) {
	obj := map[string]interface{}{
		"foo": "bar",
		"baz": "qux",
	}

	encoded, err := Encode(obj)
	if err != nil {
		t.Errorf("error encoding: %s", err)
	}

	decoded := map[string]interface{}{}
	err = Decode(encoded, &decoded)
	if err != nil {
		t.Errorf("error decoding: %s", err)
	}

	if !reflect.DeepEqual(obj, decoded) {
		t.Errorf("expected %v, got %v", obj, decoded)
	}
}

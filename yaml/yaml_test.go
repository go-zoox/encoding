package yaml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestYAML(t *testing.T) {
	obj := map[string]interface{}{
		"foo": "bar",
		"baz": "qux",
	}

	enc := New()
	encoded, err := enc.Encode(obj)
	if err != nil {
		t.Errorf("error encoding: %s", err)
	}

	decoded := map[string]interface{}{}
	err = enc.Decode(encoded, &decoded)
	if err != nil {
		t.Errorf("error decoding: %s", err)
	}

	if !reflect.DeepEqual(obj, decoded) {
		t.Errorf("expected %v, got %v", obj, decoded)
	}

	fmt.Println("obj:", obj)
	fmt.Println("decoded:", decoded)
}

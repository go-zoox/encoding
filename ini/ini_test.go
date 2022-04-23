package json

import (
	"reflect"
	"testing"
)

func TestINI(t *testing.T) {
	type Config struct {
		Foo string `ini:"foo"`
		Baz string `ini:"baz"`
	}
	obj := Config{
		Foo: "bar",
		Baz: "qux",
	}

	encoded, err := Encode(&obj)
	if err != nil {
		t.Errorf("error encoding: %s", err)
	}

	// fmt.Println("encoded:", string(encoded))

	var decoded Config
	err = Decode(encoded, &decoded)
	if err != nil {
		t.Errorf("error decoding: %s", err)
	}

	// fmt.Println("decoded:", decoded)

	if !reflect.DeepEqual(obj, decoded) {
		t.Errorf("expected %v, got %v", obj, decoded)
	}
}

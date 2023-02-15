package credentials

import (
	"testing"

	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/testify"
)

func TestCredentialsSimpleUsernameAndPassword(t *testing.T) {
	raw := "https://xxx:yyy@github.com"
	decoded, err := Decode([]byte(raw))
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, "xxx", decoded["github.com"].User.Username())
	password, _ := decoded["github.com"].User.Password()
	testify.Equal(t, "yyy", password)

	encoded, err := Encode(decoded)
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, raw, string(encoded))
}

func TestCredentialsUsernameAndPasswordNeedEscape(t *testing.T) {
	raw := "http://zero%40example.com:xa8J%5EN43&&8%25G@10.0.0.1:8888"
	decoded, err := Decode([]byte(raw))
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, "zero@example.com", decoded["10.0.0.1:8888"].User.Username())
	password, _ := decoded["10.0.0.1:8888"].User.Password()
	testify.Equal(t, "xa8J^N43&&8%G", password)

	encoded, err := Encode(decoded)
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, raw, string(encoded))
}

func TestCredentialsMultipleLines(t *testing.T) {
	raw := `https://xxx:yyy@github.com
http://zero:zero2@10.0.0.1:8888
http://zero:zero3@10.0.0.1:8888
http://zero:zero4@10.0.0.1:8888
http://zero:zero5@example.com`

	decoded, err := Decode([]byte(raw))
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(decoded)
	for _, item := range decoded {
		username := item.User.Username()
		password, _ := item.User.Password()
		fmt.Println(username, password)
	}

	testify.Equal(t, "zero", decoded["10.0.0.1:8888"].User.Username())
	password, _ := decoded["10.0.0.1:8888"].User.Password()
	testify.Equal(t, "zero4", password)

	encoded, err := Encode(decoded)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(encoded))

	testify.Equal(t, "https://xxx:yyy@github.com\nhttp://zero:zero4@10.0.0.1:8888\nhttp://zero:zero5@example.com", string(encoded))
}

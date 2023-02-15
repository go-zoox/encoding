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

	testify.Equal(t, "xxx", decoded["https://github.com:443"].Username)
	testify.Equal(t, "yyy", decoded["https://github.com:443"].Password)

	encoded, err := Encode(decoded)
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, raw, string(encoded))
}

func TestCredentialsUsernameAndPasswordNeedEscape(t *testing.T) {
	raw := "http://zero%40example.com:xa8J%5EN43%26%268%25G@10.0.0.1%3A8888"
	decoded, err := Decode([]byte(raw))
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, "zero@example.com", decoded["http://10.0.0.1:8888"].Username)
	testify.Equal(t, "xa8J^N43&&8%G", decoded["http://10.0.0.1:8888"].Password)

	encoded, err := Encode(decoded)
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, raw, string(encoded))
}

func TestCredentialsMultipleLines(t *testing.T) {
	raw := `https://xxx:yyy@github.com
http://zero:zero2@10.0.0.1%3A8888
http://zero:zero3@10.0.0.1%3A8888
http://zero:zero4@10.0.0.1%3A8888
http://zero:zero5@example.com`

	decoded, err := Decode([]byte(raw))
	if err != nil {
		t.Fatal(err)
	}

	fmt.PrintJSON(decoded)
	for _, item := range decoded {
		username := item.Username
		password := item.Password
		fmt.Println(username, password)
	}

	testify.Equal(t, "zero", decoded["http://10.0.0.1:8888"].Username)
	testify.Equal(t, "zero4", decoded["http://10.0.0.1:8888"].Password)

	encoded, err := Encode(decoded)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(encoded))

	testify.Equal(t, "https://xxx:yyy@github.com\nhttp://zero:zero5@example.com\nhttp://zero:zero4@10.0.0.1%3A8888", string(encoded))
}

package credentials

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestItemCase01(t *testing.T) {
	raw := "https://username:password@github.com"
	item, err := ParseItem(raw)
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, "https", item.Protocol)
	testify.Equal(t, "username", item.Username)
	testify.Equal(t, "password", item.Password)
	testify.Equal(t, "github.com", item.Host)
	testify.Equal(t, 443, item.Port)

	testify.Equal(t, raw, item.String())
}

func TestItemCase02(t *testing.T) {
	raw := "http://username:password@127.0.0.1%3A8888"
	item, err := ParseItem(raw)
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, "http", item.Protocol)
	testify.Equal(t, "username", item.Username)
	testify.Equal(t, "password", item.Password)
	testify.Equal(t, "127.0.0.1", item.Host)
	testify.Equal(t, 8888, item.Port)

	testify.Equal(t, raw, item.String())
}

func TestItemCase03(t *testing.T) {
	raw := "http://zero%40example.com:xa8J%5EN43%26%268%25G@10.0.0.1%3A8888"
	item, err := ParseItem(raw)
	if err != nil {
		t.Fatal(err)
	}

	testify.Equal(t, "http", item.Protocol)
	testify.Equal(t, "zero@example.com", item.Username)
	testify.Equal(t, "xa8J^N43&&8%G", item.Password)
	testify.Equal(t, "10.0.0.1", item.Host)
	testify.Equal(t, 8888, item.Port)

	testify.Equal(t, raw, item.String())
}

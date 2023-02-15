package credentials

import (
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/go-zoox/core-utils/object"
)

type Credentials interface {
	Encode(v map[string]*Item) ([]byte, error)
	Decode(raw []byte) (map[string]*Item, error)
}

type Item = url.URL

type credentials struct {
}

func New() Credentials {
	return &credentials{}
}

func (c credentials) Encode(v map[string]*Item) ([]byte, error) {
	lines := []string{}
	items := object.Values(v)

	sort.Slice(items, func(i, j int) bool {
		return strings.Compare(items[i].Host, items[j].Host) > -1
	})

	for _, one := range v {
		// @TODO git credentials store encode escape characters not standard
		// https://en.wikipedia.org/wiki/URL_encoding
		lines = append(lines, one.String())
	}

	return []byte(strings.Join(lines, "\n")), nil
}

func (c credentials) Decode(raw []byte) (map[string]*Item, error) {
	v := make(map[string]*Item, 0)

	lines := strings.Split(string(raw), "\n")
	for _, one := range lines {
		one = strings.TrimSpace(one)

		// empty line
		if len(one) == 0 {
			continue
		}

		// comment
		if one[0] == '#' {
			continue
		}

		item, err := url.Parse(one)
		if err != nil {
			return nil, fmt.Errorf("invalid credentials(%s): %v", one, err)
		}

		v[item.Host] = item
	}

	return v, nil
}

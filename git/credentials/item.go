package credentials

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Item struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	//
	raw string `json:"raw"`
}

func (i *Item) Parse(raw string) error {
	i.raw = raw

	// https://username:password@github.com
	// http://username:password@127.0.0.1:2080
	// http://username:password@127.0.0.1%3A2080

	rest := raw
	parts := strings.Split(rest, "://")
	if len(parts) != 2 {
		return fmt.Errorf("cannot get protocol")
	}

	// protocol
	i.Protocol = parts[0]
	rest = parts[1]

	// username + password
	parts = strings.Split(rest, "@")
	if len(parts) != 2 {
		return fmt.Errorf("cannot get username and password")
	}
	if err := i.parseAuth(parts[0]); err != nil {
		return fmt.Errorf("failed to parse auth: %v", err)
	}
	rest = parts[1]

	// host + port
	if err := i.parseHost(rest); err != nil {
		return fmt.Errorf("failed to parse host + port: %v", err)
	}

	return nil
}

func (i *Item) String() string {
	host := i.generateHost()
	auth := i.generateAuth()
	return fmt.Sprintf("%s://%s@%s", i.Protocol, auth, host)
}

func (i *Item) Raw() string {
	return i.raw
}

func (i *Item) escape(s string) string {
	return url.QueryEscape(s)
}

func (i *Item) unescape(s string) (string, error) {
	return url.QueryUnescape(s)
}

func (i *Item) generateAuth() string {
	Username := i.escape(i.Username)
	Password := i.escape(i.Password)
	return fmt.Sprintf("%s:%s", Username, Password)
}

func (i *Item) parseAuth(auth string) (err error) {
	parts := strings.Split(auth, ":")
	if len(parts) != 2 {
		return fmt.Errorf("invalid auth(%s)", auth)
	}

	i.Username, err = i.unescape(parts[0])
	if err != nil {
		return fmt.Errorf("failed to parse username(%s)", parts[0])
	}

	i.Password, err = i.unescape(parts[1])
	if err != nil {
		return fmt.Errorf("failed to parse username(%s)", parts[1])
	}

	return nil
}

func (i *Item) generateHost() string {
	if i.Port == 80 || i.Port == 443 {
		return i.escape(i.Host)
	}

	return i.escape(fmt.Sprintf("%s:%d", i.Host, i.Port))
}

func (i *Item) parseHost(host string) error {
	hostx, err := i.unescape(host)
	if err != nil {
		return fmt.Errorf("failed to unescape host(%s): %v", host, err)
	}

	parts := strings.Split(hostx, ":")
	if len(parts) > 2 {
		return fmt.Errorf("invalid host(%s)", hostx)
	}

	i.Host, err = i.unescape(parts[0])
	if err != nil {
		return fmt.Errorf("failed to parse host(%s)", parts[0])
	}

	if len(parts) == 2 {
		port, err := i.unescape(parts[1])
		if err != nil {
			return fmt.Errorf("failed to parse port(%s)", parts[1])
		}

		i.Port, err = strconv.Atoi(port)
		if err != nil {
			return fmt.Errorf("failed to parse port: %v", err)
		}
	}

	if i.Port == 0 {
		switch i.Protocol {
		case "http":
			i.Port = 80
		case "https":
			i.Port = 443
		}
	}

	return nil
}

func ParseItem(raw string) (*Item, error) {
	var item Item
	if err := item.Parse(raw); err != nil {
		return nil, err
	}

	return &item, nil
}

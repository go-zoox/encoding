# Encoding - Unmashal/Marsha of JSON/YAML/TOML 

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/encoding)](https://pkg.go.dev/github.com/go-zoox/encoding)
[![Build Status](https://github.com/go-zoox/encoding/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/encoding/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/encoding)](https://goreportcard.com/report/github.com/go-zoox/encoding)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/encoding/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/encoding?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/encoding.svg)](https://github.com/go-zoox/encoding/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/encoding.svg?label=Release)](https://github.com/go-zoox/encoding/releases)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/encoding
```

## Getting Started

```go
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
}
```

## Inspired by

## License
GoZoox is released under the [MIT License](./LICENSE).
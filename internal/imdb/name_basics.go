package imdb

import (
	"errors"
	"strconv"
)

func equalFields(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func parseOptionalInt(v string) (*int, error) {
	if v == `\N` {
		return nil, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return nil, err
	}
	return &n, nil
}

func ioEOF() error {
	return errors.New("EOF")
}

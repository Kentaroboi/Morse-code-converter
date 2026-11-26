package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert here we try to understand what we get - morse or simple text,
// and then convert it to the other form. Text we make to morse, morse we make to text.
func Convert(data string) (string, error) {
	data = strings.TrimSpace(data)
	if data == "" {
		return "", errors.New("empty input")
	}

	if isMorse(data) {
		return morse.ToText(data), nil
	}
	return morse.ToMorse(data), nil
}

// isMorse check that string have only morse symbols - dots, dashes and spaces and newlines.
func isMorse(data string) bool {
	for _, r := range data {
		if r == '.' || r == '-' || unicode.IsSpace(r) {
			continue
		}
		return false
	}
	return true
}

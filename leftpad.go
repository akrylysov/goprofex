package main

import (
	"strings"
)

func leftpad(s string, length int, char rune) string {
	if len(s) < length {
		return strings.Repeat(string(char), length-len(s)) + s
	}
	return s
}

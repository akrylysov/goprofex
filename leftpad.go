package main

func leftpad(s string, length int, char rune) string {
	for len(s) < length {
		s = string(char) + s
	}
	return s
}

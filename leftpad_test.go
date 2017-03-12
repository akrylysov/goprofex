package main

import "testing"

func benchmarkLeftpad(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		leftpad("test", n, '*')
	}
}

func BenchmarkLeftpad10(b *testing.B) {
	benchmarkLeftpad(b, 10)
}

func BenchmarkLeftpad50(b *testing.B) {
	benchmarkLeftpad(b, 50)
}

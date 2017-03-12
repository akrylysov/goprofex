package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkTimedHandler(b *testing.B) {
	r := httptest.NewRequest("GET", "/test", nil)
	h := timedHandler("test", func(w http.ResponseWriter, r *http.Request) {})
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		h(w, r)
	}
}

func BenchmarkLeftpadHandler(b *testing.B) {
	r := httptest.NewRequest("GET", "/v1/leftpad/?str=test&len=50&chr=*", nil)
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		leftpadHandler(w, r)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type leftpadResponse struct {
	Str string `json:"str"`
}

func timedHandler(name string, nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		nextFunc(w, r)
		elapsed := time.Since(start)
		statsd.Timing(fmt.Sprintf("request.%s.timing", name), elapsed)
		statsd.Incr(fmt.Sprintf("request.%s.count", name))
		log.Printf("%s request took %v", name, elapsed)
	}
}

func leftpadHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("url %s\tip %s\tua %s", r.RequestURI, r.RemoteAddr, r.UserAgent())
	str := r.FormValue("str")
	length, err := strconv.Atoi(r.FormValue("len"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	statsd.Histogram("request.str.len", float64(length))
	chr := ' '
	if len(r.FormValue("chr")) > 0 {
		chr = []rune(r.FormValue("chr"))[0]
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := leftpadResponse{Str: leftpad(str, length, chr)}
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

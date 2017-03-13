package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var statsd = &StatsD{Namespace: "leftpad", SampleRate: 0.5}

func init() {
	var f, err = ioutil.TempFile("", "leftpad.log")
	if err != nil {
		panic(err)
	}
	log.SetOutput(bufio.NewWriterSize(f, 1024*16))
}

func main() {
	http.HandleFunc("/v1/leftpad/", timedHandler("leftpad", leftpadHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

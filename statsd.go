package main

import (
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type StatsD struct {
	Namespace  string
	SampleRate float64
}

var reservedReplacer = strings.NewReplacer(":", "_", "|", "_", "@", "_")

func (s *StatsD) Send(stat string, kind string, delta float64) {
	buf := bytes.Buffer{}
	buf.WriteString(s.Namespace)
	buf.WriteByte('.')
	buf.WriteString(reservedReplacer.Replace(stat))
	buf.WriteByte(':')
	buf.Write(strconv.AppendFloat(make([]byte, 0, 24), delta, 'f', -1, 64))
	buf.WriteByte('|')
	buf.WriteString(kind)
	if s.SampleRate != 0 && s.SampleRate < 1 {
		buf.WriteString("|@")
		buf.Write(strconv.AppendFloat(make([]byte, 0, 24), s.SampleRate, 'f', -1, 64))
	}
	buf.WriteTo(ioutil.Discard) // TODO: Write to a socket
}

func (s *StatsD) Incr(stat string) {
	s.Send(stat, "c", 1)
}

func (s *StatsD) Histogram(stat string, value float64) {
	s.Send(stat, "h", value)
}

func (s *StatsD) Timing(stat string, value time.Duration) {
	s.Send(stat, "ms", value.Seconds()*1000)
}

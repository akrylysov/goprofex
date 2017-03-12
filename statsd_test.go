package main

import "testing"

func BenchmarkStatsD(b *testing.B) {
	statsd := StatsD{Namespace: "namespace", SampleRate: 0.5}
	for i := 0; i < b.N; i++ {
		statsd.Incr("test")
	}
}

package jay_test

import (
	"github.com/speedyhoon/jay"
	"testing"
)

var s string
var l int
var ok bool
var y = make([]byte, 19)
var z = []byte("\022octopus camouflage")

func BenchmarkReadString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, l, ok = jay.ReadString(z)
	}
}

func BenchmarkReadStringPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l, ok = jay.ReadStringPtr(z, &s)
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.String(y, "octopus camouflage", 18)
	}
}

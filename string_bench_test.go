package jay_test

import (
	"github.com/speedyhoon/jay"
	"testing"
)

var str1 string
var y = make([]byte, 19)
var z = []byte("\022octopus camouflage")

func BenchmarkReadString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str1, integer, bool1 = jay.ReadString(z)
	}
}

func BenchmarkReadStringPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		integer, bool1 = jay.ReadStringPtr(z, &str1)
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jay.String(y, "octopus camouflage", 18)
	}
}

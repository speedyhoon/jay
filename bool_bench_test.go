package jay

import (
	"testing"
)

var m, n, o, p, q, r, s, t bool

func BenchmarkReadBool1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m = ReadBool1(byte(i))
	}
}
func BenchmarkReadBool2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m, n = ReadBool2(byte(i))
	}
}
func BenchmarkReadBool3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m, n, o = ReadBool3(byte(i))
	}
}
func BenchmarkReadBool4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m, n, o, p = ReadBool4(byte(i))
	}
}
func BenchmarkReadBool5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m, n, o, p, q = ReadBool5(byte(i))
	}
}
func BenchmarkReadBool6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m, n, o, p, q, r = ReadBool6(byte(i))
	}
}
func BenchmarkReadBool7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m, n, o, p, q, r, s = ReadBool7(byte(i))
	}
}
func BenchmarkReadBool8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m, n, o, p, q, r, s, t = ReadBool8(byte(i))
	}
}

var y byte

func BenchmarkBool1True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = Bool1(true)
	}
}
func BenchmarkBool1False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = Bool1(false)
	}
}
func BenchmarkBool2True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = Bool2(true, true)
	}
}
func BenchmarkBool2False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = Bool2(false, false)
	}
}
func BenchmarkBool3True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = Bool3(true, true, true)
	}
}
func BenchmarkBool3False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = Bool3(false, false, false)
	}
}
func BenchmarkBool4True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = Bool4(true, true, true, true)
	}
}
func BenchmarkBool4False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = Bool4(false, false, false, false)
	}
}
func BenchmarkBool5True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = Bool5(true, true, true, true, true)
	}
}
func BenchmarkBool5False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = Bool5(false, false, false, false, false)
	}
}
func BenchmarkBool6True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = Bool6(true, true, true, true, true, true)
	}
}
func BenchmarkBool6False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = Bool6(false, false, false, false, false, false)
	}
}
func BenchmarkBool7True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = Bool7(true, true, true, true, true, true, true)
	}
}
func BenchmarkBool7False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = Bool7(false, false, false, false, false, false, false)
	}
}
func BenchmarkBool8True(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = Bool8(true, true, true, true, true, true, true, true)
	}
}
func BenchmarkBool8False(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y = Bool8(false, false, false, false, false, false, false, false)
	}
}

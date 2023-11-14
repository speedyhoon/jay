package main

import "testing"

var car = Car{ID: 42, Auto: true, Name: "Hello", Gearbox: Gearbox{
	Gears:        4,
	Reverse:      1,
	Sequential:   true,
	Automatic:    false,
	Model:        "H23w",
	Manufacturer: "Hewland",
}}
var src []byte

func BenchmarkMarshalJ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src = car.MarshalJ()
	}
}

func BenchmarkMarshalJBuf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src = car.MarshalJ2()
	}
}
func BenchmarkMarshalJ2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = car.MarshalJ()
	}
}
func BenchmarkMarshalK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = car.MarshalK()
	}
}
func BenchmarkMarshalL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = car.MarshalL()
	}
}

var ll int

func BenchmarkLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ll = len(src)
	}
}
func BenchmarkFastest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fastest(i)
	}
}
func Fastest(i int) {
	// nothing to do here
	_ = i + 1
}

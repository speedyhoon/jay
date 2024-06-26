package jay_test

import (
	"github.com/speedyhoon/jay"
	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"testing"
)

func TestFloat32s(t *testing.T) {
	b := make([]byte, 1)
	t.Run("nil", func(t *testing.T) {
		jay.WriteFloat32s(b, nil, 0)
		assert.Equal(t, []float32(nil), jay.ReadFloat32s(b, 0))
	})

	t.Run("empty slice", func(t *testing.T) {
		jay.WriteFloat32s(b, []float32{}, 0)
		assert.Equal(t, []float32(nil), jay.ReadFloat32s(b, 0))
	})

	f := []float32{rando.Float32()}
	b = make([]byte, 4)
	t.Run("one", func(t *testing.T) {
		jay.WriteFloat32s(b, f, len(f))
		assert.Equal(t, f, jay.ReadFloat32s(b, len(f)))
	})

	f = []float32{math.MaxFloat32}
	t.Run("max", func(t *testing.T) {
		jay.WriteFloat32s(b, f, len(f))
		assert.Equal(t, f, jay.ReadFloat32s(b, len(f)))
	})

	f = []float32{math.SmallestNonzeroFloat32}
	t.Run("smallest", func(t *testing.T) {
		jay.WriteFloat32s(b, f, len(f))
		assert.Equal(t, f, jay.ReadFloat32s(b, len(f)))
	})

	for i := 0; i < 10; i++ {
		f = rando.Float32s()
		b = make([]byte, len(f)*4)
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			jay.WriteFloat32s(b, f, len(f))
			assert.Equal(t, f, jay.ReadFloat32s(b, len(f)))
		})
	}
}

func TestFloat64s(t *testing.T) {
	b := make([]byte, 1)
	t.Run("nil", func(t *testing.T) {
		jay.WriteFloat64s(b, nil, 0)
		assert.Equal(t, []float64(nil), jay.ReadFloat64s(b, 0))
	})

	t.Run("empty slice", func(t *testing.T) {
		jay.WriteFloat64s(b, []float64{}, 0)
		assert.Equal(t, []float64(nil), jay.ReadFloat64s(b, 0))
	})

	f := []float64{rando.Float64()}
	b = make([]byte, 8)
	t.Run("one", func(t *testing.T) {
		jay.WriteFloat64s(b, f, len(f))
		assert.Equal(t, f, jay.ReadFloat64s(b, len(f)))
	})

	f = []float64{math.MaxFloat64}
	t.Run("max", func(t *testing.T) {
		jay.WriteFloat64s(b, f, len(f))
		assert.Equal(t, f, jay.ReadFloat64s(b, len(f)))
	})

	f = []float64{math.SmallestNonzeroFloat64}
	t.Run("smallest", func(t *testing.T) {
		jay.WriteFloat64s(b, f, len(f))
		assert.Equal(t, f, jay.ReadFloat64s(b, len(f)))
	})

	for i := 0; i < 10; i++ {
		f = rando.Float64s()
		b = make([]byte, len(f)*8)
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			jay.WriteFloat64s(b, f, len(f))
			assert.Equal(t, f, jay.ReadFloat64s(b, len(f)))
		})
	}
}

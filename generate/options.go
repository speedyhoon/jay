package generate

import (
	"fmt"
	"math"
)

const (
	Auto = MaxSize(iota * 32)
	Bit32
	Bit64
)

type MaxSize uint8

type Option struct {
	MaxIntSize  MaxSize
	MaxUintSize MaxSize

	Is32bit bool

	// MaxDefaultStrSize limits all strings to be within this length if a field tag is NOT present.
	// Minimum:	1 (1 byte),
	// Default: MaxUint8 (255 bytes),
	// Maximum: MaxUint24 (16 Megabytes).
	// To override the default for a field use: `j:"max:4030"` for 4,030 bytes.
	// The smallest value is the most optimal for performance.
	MaxDefaultStrSize uint
	strSizeOfDefault  uint8

	// MaxUint16 = 64 kilobytes,
	// MaxUint24 = 16 Megabytes,
	//4 Gigabytes
	//	// MaxUint8 = 255 bytes (default),
	//	// MaxUint32 = 4 Gigabytes (maximum).

	// Should integers be fixed to 4 or 8 bytes or vary in length depending on the value provided.
	FixedIntSize bool

	// Whether unsigned integers be a fixed length (4 bytes 32-bit, or 8 bytes 64-bit) or vary in length depending on the value provided.
	// True = Highest CPU serialization/deserialization throughput,
	// false = Least bandwidth used.
	FixedUintSize bool

	// TODO add option to check if a struct or map is nil/empty by appending an extra null byte \0x0 ?
	// If the null byte wasn't there - how would the Read functions know if there was an unexpected
	// end of buffer vs the struct/map was empty?

	Verbose      bool
	IncludeTests bool
}

func (m *MaxSize) Set(value *uint) error {
	if *value > 8 {
		return fmt.Errorf("value is greater than 8. Expected 1 - 8, got %d", *value)
	}

	*m = MaxSize(*value)
	return nil
}

func LoadOptions(opts ...Option) (o Option) {
	if len(opts) >= 1 {
		o = opts[0]
	}

	if o.MaxDefaultStrSize == 0 {
		o.MaxDefaultStrSize = MaxUint8
	} else if o.MaxDefaultStrSize > MaxUint24 {
		o.MaxDefaultStrSize = MaxUint24
	}
	o.strSizeOfDefault = bytesRequired(o.MaxDefaultStrSize)

	if o.MaxIntSize == Auto || o.MaxIntSize > Bit32 && o.MaxIntSize < Bit64 {
		o.MaxIntSize = 32 << (^uint(0) >> 63) // 32 or 64
		return
	}

	if o.MaxIntSize > Bit64 {
		o.MaxIntSize = Bit64
		return
	}

	if o.MaxIntSize < Bit32 {
		o.MaxIntSize = Bit32
	}
	return
}

// bytesRequired returns how many bytes are required to represent an unsigned integer.
func bytesRequired(input uint) uint8 {
	if input <= 1 {
		return uint8(input)
	}
	return uint8(math.Ceil(LogBaseX(256, float64(input+1))))
}

func LogBaseX(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}

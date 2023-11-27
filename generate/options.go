package generate

import (
	"github.com/speedyhoon/jay"
)

const (
	Auto uint8 = iota * 32
	Bit32
	Bit64
)

type Option struct {
	MaxIntegerSize uint8

	Is32bit bool

	// MaxDefaultStrSize limits all strings to be within this length if a field tag is NOT present.
	// Minimum:	1 (1 byte),
	// Default: MaxUint8 (255 bytes),
	// Maximum: MaxUint24 (16 Megabytes).
	// To override the default for a field use: `j:"max:4030"` for 4,030 bytes.
	// The smallest value is the most optimal for performance.
	MaxDefaultStrSize uint32

	// MaxUint16 = 64 kilobytes,
	// MaxUint24 = 16 Megabytes,
	//4 Gigabytes
	//	// MaxUint8 = 255 bytes (default),
	//	// MaxUint32 = 4 Gigabytes (maximum).

	// Should integers be fixed to 4 or 8 bytes or vary in length depending on the value provided.
	FixedIntSize bool

	// Whether unsigned integers be a fixed length (4 bytes 32-bit, or 8 bytes 64-bit) or vary in length depending on the value provided.
	// true = Highest CPU serialization/deserialization throughput,
	// false = Least bandwidth used.
	FixedUintSize bool

	// TODO add option to check if a struct or map is nil/empty by appending an extra null byte \0x0 ?
	// If the null byte wasn't there - how would the Read functions know if there was an unexpected
	// end of buffer vs the struct/map was empty?
}

func LoadOptions(opts ...Option) (o Option) {
	if len(opts) >= 1 {
		o = opts[0]
	}

	if o.MaxDefaultStrSize == 0 {
		o.MaxDefaultStrSize = jay.MaxUint24
	}

	if o.MaxIntegerSize == Auto || o.MaxIntegerSize > Bit32 && o.MaxIntegerSize < Bit64 {
		o.MaxIntegerSize = 32 << (^uint(0) >> 63) // 32 or 64
		return
	}

	if o.MaxIntegerSize > Bit64 {
		o.MaxIntegerSize = Bit64
		return
	}

	if o.MaxIntegerSize < Bit32 {
		o.MaxIntegerSize = Bit32
	}
	return
}
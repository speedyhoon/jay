package jay

const (
	Auto uint8 = iota * 32
	Bit32
	Bit64
)

type Option struct {
	MaxIntegerSize uint8

	// Should integers be fixed to 4 or 8 bytes or vary in length depending on the value provided.
	FixedIntSize bool

	// Should unsigned integers be fixed to 4 or 8 bytes or vary in length depending on the value provided.
	FixedUintSize bool
}

func LoadOptions(opts ...Option) (o Option) {
	if len(opts) >= 1 {
		o = opts[0]
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
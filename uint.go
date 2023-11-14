package jay

const (
	MaxUint8  = 1<<8 - 1  // 255
	MaxUint16 = 1<<16 - 1 // 65535
	MaxUint24 = 1<<24 - 1 // 16777215
	MaxUint32 = 1<<32 - 1 // 4294967295
	MaxUint40 = 1<<40 - 1 // 1099511627775
	MaxUint48 = 1<<48 - 1 // 281474976710655
	MaxUint56 = 1<<56 - 1 // 72057594037927935
	//MaxUint64 = 1<<64 - 1 // 18446744073709551615
)

func LenUint(i uint) uint8 {
	switch {
	case i == 0:
		return 1
	case i <= MaxUint8:
		return 2
	case i <= MaxUint16:
		return 3
	case i <= MaxUint24:
		return 4
	case i <= MaxUint32:
		return 5
	case i <= MaxUint40:
		return 6
	case i <= MaxUint48:
		return 7
	case i <= MaxUint56:
		return 8
	default:
		return 9
	}
}

func WriteUintBytes(y []byte, u uint, length uint8) {
	y[0] = length
	switch length {
	// case 1: No further processing required.
	case 2:
		y[1] = byte(u)
	case 3:
		y[1] = byte(u)
		y[2] = byte(u >> 8)
	case 4:
		y[1] = byte(u)
		y[2] = byte(u >> 8)
		y[3] = byte(u >> 16)
	case 5:
		y[1] = byte(u)
		y[2] = byte(u >> 8)
		y[3] = byte(u >> 16)
		y[4] = byte(u >> 24)
	case 6:
		y[1] = byte(u)
		y[2] = byte(u >> 8)
		y[3] = byte(u >> 16)
		y[4] = byte(u >> 24)
		y[5] = byte(u >> 32)
	case 7:
		y[1] = byte(u)
		y[2] = byte(u >> 8)
		y[3] = byte(u >> 16)
		y[4] = byte(u >> 24)
		y[5] = byte(u >> 32)
		y[6] = byte(u >> 40)
	case 8:
		y[1] = byte(u)
		y[2] = byte(u >> 8)
		y[3] = byte(u >> 16)
		y[4] = byte(u >> 24)
		y[5] = byte(u >> 32)
		y[6] = byte(u >> 40)
		y[7] = byte(u >> 48)
	case 9:
		y[1] = byte(u)
		y[2] = byte(u >> 8)
		y[3] = byte(u >> 16)
		y[4] = byte(u >> 24)
		y[5] = byte(u >> 32)
		y[6] = byte(u >> 40)
		y[7] = byte(u >> 48)
		y[8] = byte(u >> 56)
	}
}

func WriteUint64Bytes(y []byte, u uint64) {
	y[0], y[1], y[2], y[3], y[4], y[5], y[6], y[7] = byte(u), byte(u>>8), byte(u>>16), byte(u>>24), byte(u>>32), byte(u>>40), byte(u>>48), byte(u>>56)
}

func WriteUint32Bytes(y []byte, u uint32) {
	y[0], y[1], y[2], y[3] = byte(u), byte(u>>8), byte(u>>16), byte(u>>24)
}

func WriteUint16Bytes(y []byte, u uint16) {
	y[0], y[1] = byte(u), byte(u>>8)
}

func WriteUint56Bytes(b []byte, i uint64) {
	b[0] = byte(i)
	b[1] = byte(i >> 8)
	b[2] = byte(i >> 16)
	b[3] = byte(i >> 24)
	b[4] = byte(i >> 32)
	b[5] = byte(i >> 40)
	b[6] = byte(i >> 48)
}

func WriteUint48Bytes(b []byte, i uint64) {
	b[0] = byte(i)
	b[1] = byte(i >> 8)
	b[2] = byte(i >> 16)
	b[3] = byte(i >> 24)
	b[4] = byte(i >> 32)
	b[5] = byte(i >> 40)
}
func WriteUint40Bytes(b []byte, i uint64) {
	b[0] = byte(i)
	b[1] = byte(i >> 8)
	b[2] = byte(i >> 16)
	b[3] = byte(i >> 24)
	b[4] = byte(i >> 32)
}
func WriteUint24Bytes(b []byte, i uint64) {
	b[0] = byte(i)
	b[1] = byte(i >> 8)
	b[2] = byte(i >> 16)
}

// DEPRECATED
func LenUintDEPRECATED(i uint) uint8 {
	switch {
	case i == 0:
		return 0
	case i <= MaxUint8:
		return 1
	case i <= MaxUint16:
		return 2
	case i <= MaxUint24:
		return 3
	case i <= MaxUint32:
		return 4
	case i <= MaxUint40:
		return 5
	case i <= MaxUint48:
		return 6
	case i <= MaxUint56:
		return 7
	default:
		return 8
	}
}

// DEPRECATED
func WriteUintBytesDEPRECATED(b []byte, i uint) int {
	b[0] = LenUintDEPRECATED(i)
	switch b[0] {
	default:
		// Not required, value is zero.
		return 1
	case 1:
		b[1] = byte(i)
		return 2
	case 2:
		b[1] = byte(i)
		b[2] = byte(i >> 8)
		return 3
	case 3:
		b[1] = byte(i)
		b[2] = byte(i >> 8)
		b[3] = byte(i >> 16)
		return 4
	case 4:
		b[1] = byte(i)
		b[2] = byte(i >> 8)
		b[3] = byte(i >> 16)
		b[4] = byte(i >> 24)
		return 5
	case 5:
		b[1] = byte(i)
		b[2] = byte(i >> 8)
		b[3] = byte(i >> 16)
		b[4] = byte(i >> 24)
		b[5] = byte(i >> 32)
		return 6
	case 6:
		b[1] = byte(i)
		b[2] = byte(i >> 8)
		b[3] = byte(i >> 16)
		b[4] = byte(i >> 24)
		b[5] = byte(i >> 32)
		b[6] = byte(i >> 40)
		return 7
	case 7:
		b[1] = byte(i)
		b[2] = byte(i >> 8)
		b[3] = byte(i >> 16)
		b[4] = byte(i >> 24)
		b[5] = byte(i >> 32)
		b[6] = byte(i >> 40)
		b[7] = byte(i >> 48)
		return 8
	case 8:
		b[1] = byte(i)
		b[2] = byte(i >> 8)
		b[3] = byte(i >> 16)
		b[4] = byte(i >> 24)
		b[5] = byte(i >> 32)
		b[6] = byte(i >> 40)
		b[7] = byte(i >> 48)
		b[8] = byte(i >> 56)
		return 9
	}
}

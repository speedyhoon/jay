package jay

// Bit shift masks.
const _24, _40, _48, _56 = 24, 40, 48, 56

// ReadInt64 ...
func ReadInt64(b []byte) int64 {
	return int64(b[0]) | int64(b[1])<<_8 | int64(b[2])<<_16 | int64(b[3])<<_24 |
		int64(b[4])<<_32 | int64(b[5])<<_40 | int64(b[6])<<_48 | int64(b[7])<<_56
}

// ReadIntArch64 ...
func ReadIntArch64(b []byte) int {
	return int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16 | int(b[3])<<_24 |
		int(b[4])<<_32 | int(b[5])<<_40 | int(b[6])<<_48 | int(b[7])<<_56
}

// ReadIntArch32 ...
func ReadIntArch32(b []byte) int {
	return int(b[0]) | int(b[1])<<_8 | int(b[2])<<_16 | int(b[3])<<_24
}

// ReadInt32 ...
func ReadInt32(b []byte) int32 {
	return int32(b[0]) | int32(b[1])<<_8 | int32(b[2])<<_16 | int32(b[3])<<_24
}

// ReadInt16 ...
func ReadInt16(b []byte) int16 {
	return int16(b[0]) | int16(b[1])<<_8
}

// WriteInt64 ...
func WriteInt64(b []byte, i int64) {
	b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32), byte(i>>_40), byte(i>>_48), byte(i>>_56)
}

// WriteIntArch64 ...
func WriteIntArch64(b []byte, i int) {
	b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24), byte(i>>_32), byte(i>>_40), byte(i>>_48), byte(i>>_56)
}

// WriteIntArch32 ...
func WriteIntArch32(b []byte, i int) {
	b[0], b[1], b[2], b[3] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24)
}

// WriteInt32 ...
func WriteInt32(b []byte, i int32) {
	b[0], b[1], b[2], b[3] = byte(i), byte(i>>_8), byte(i>>_16), byte(i>>_24)
}

// WriteInt16 ...
func WriteInt16(b []byte, i int16) {
	b[0], b[1] = byte(i), byte(i>>_8)
}

package generate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalBoolsInline(t *testing.T) {
	assert.Equal(t, "", unmarshalBoolsInline("z", 55, 0), "zero")
	assert.Equal(t, "b[55] >= 128", unmarshalBoolsInline("b", 55, 1), "one")
	assert.Equal(t, "e[67] >= 128, e[67]&64 == 64", unmarshalBoolsInline("e", 67, 2), "two")
	assert.Equal(t, "t[11] >= 128, t[11]&64 == 64, t[11]&32 == 32", unmarshalBoolsInline("t", 11, 3), "three")
	assert.Equal(t, "f[83] >= 128, f[83]&64 == 64, f[83]&32 == 32, f[83]&16 == 16", unmarshalBoolsInline("f", 83, 4), "four")
	assert.Equal(t, "q[99] >= 128, q[99]&64 == 64, q[99]&32 == 32, q[99]&16 == 16, q[99]&8 == 8", unmarshalBoolsInline("q", 99, 5), "five")
	assert.Equal(t, "u[46] >= 128, u[46]&64 == 64, u[46]&32 == 32, u[46]&16 == 16, u[46]&8 == 8, u[46]&4 == 4", unmarshalBoolsInline("u", 46, 6), "six")
	assert.Equal(t, "m[70] >= 128, m[70]&64 == 64, m[70]&32 == 32, m[70]&16 == 16, m[70]&8 == 8, m[70]&4 == 4, m[70]&2 == 2", unmarshalBoolsInline("m", 70, 7), "seven")
	assert.Equal(t, "p[5] >= 128, p[5]&64 == 64, p[5]&32 == 32, p[5]&16 == 16, p[5]&8 == 8, p[5]&4 == 4, p[5]&2 == 2, p[5]&1 == 1", unmarshalBoolsInline("p", 5, 8), "eight")
	assert.Equal(t, "", unmarshalBoolsInline("z", 55, 9), "nine")
}

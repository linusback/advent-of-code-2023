package util

import "fmt"

func StringToUint16Arr(s string) []uint16 {
	rLen := len(s) / 2
	rLen += len(s) % 2
	r := make([]uint16, rLen)
	for i := 0; i < len(s); i++ {
		if i&1 == 0 {
			r[i/2] += uint16(s[i]) << 8
		} else {
			r[i/2] += uint16(s[i])
		}
	}
	return r
}

const byteMax uint16 = 255

func Uint16ToString(src []uint16) string {
	b := make([]byte, len(src)*2)
	for i := 0; i < len(src); i++ {
		b[i*2] = byte(src[i] >> 8)
		if src[i]&byteMax != 0 {
			b[i*2+1] = byte(src[i] & byteMax)
		}
	}
	return string(b)
}

const (
	hundreds = uint16(26 * 26)
	tens     = uint16(26)
)

// change to uint16 with base 26 so AAA = 0 AND ZZZ = 17575 (26*26*26 - 1)
func toUint16(b []byte) uint16 {
	var r uint16
	if len(b) != 3 {
		panic(fmt.Sprintf("wrong length of %s", string(b)))
	}
	r += uint16(b[0]-'A') * hundreds
	r += uint16(b[1]-'A') * tens
	r += uint16(b[2] - 'A')
	return r
}

func toString(u uint16) string {
	var b [3]byte
	b[2] = byte(u%tens) + 'A'
	b[1] = byte((u%hundreds)/tens) + 'A'
	b[0] = byte(u/hundreds) + 'A'
	return string(b[:])
}

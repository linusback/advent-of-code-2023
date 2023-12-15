package util

import (
	"math/big"
)

func ValidCoordinate(x, y, lenX, lenY int64) bool {
	return x >= 0 && y >= 0 && x < lenX && y < lenY
}

func Permutate[S ~[][]K, K comparable](s S) [][]K {
	var size uint64 = 1
	for i := 0; i < len(s); i++ {
		size *= uint64(len(s[i]))
	}
	r := make([][]K, size)
	//TODO fix might be broken check day12 comment for example
	for i := 0; i < len(r); i++ {
		r[i] = make([]K, len(s))
		for j := 0; j < len(s); j++ {
			r[i][j] = s[j][i%len(s[j])]
		}
	}
	return r
}

func FindPairs[A ~[]K, K any](arr A) (res []A) {
	var (
		f    big.Int
		aLen uint64
	)
	f.MulRange(1, int64(len(arr)))
	res = make([]A, 0, len(arr))
	for i := uint64(0); i < aLen; i++ {
		for k := i + 1; k < aLen; k++ {
			e := make([]K, 2)
			e[0] = arr[i]
			e[1] = arr[k]
			res = append(res, e)
		}
	}
	return res
}

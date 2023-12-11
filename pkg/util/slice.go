package util

func ValidCoordinate(x, y, lenX, lenY int64) bool {
	return x >= 0 && y >= 0 && x < lenX && y < lenY
}

func Permutate[S ~[][]K, K comparable](s S) [][]K {
	var size uint64 = 1
	for i := 0; i < len(s); i++ {
		size *= uint64(len(s[i]))
	}
	r := make([][]K, size)
	for i := 0; i < len(r); i++ {
		r[i] = make([]K, len(s))
		for j := 0; j < len(s); j++ {
			r[i][j] = s[j][i%len(s[j])]
		}
	}
	return r
}

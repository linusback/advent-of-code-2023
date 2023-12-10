package util

import "slices"

func Unique[S ~[]K, K comparable](s S) []K {
	res := make([]K, 0, len(s))
	for i := 0; i < len(s); i++ {
		if !slices.Contains(res, s[i]) {
			res = append(res, s[i])
		}
	}
	return res
}
func FilterUnique[S ~[]K, K comparable](s S) []K {
	found := false
	for i := 0; i < len(s); i++ {
		found = false
		for j := 0; j < i; j++ {
			if s[i] == s[j] {
				found = true
				break
			}
		}
		if !found {
			continue
		}
		if i+1 < len(s) {
			//use copy to preserve order
			copy(s[i:], s[i+1:])
			// move one instead of all
			//s[i] = s[len(s)-1]
		}
		s = s[:len(s)-1]
		i--
	}
	return s
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

package util

func MapToSelfReferringArrUint16[M ~map[K][]K, K comparable](m M) ([][]uint16, map[K]uint16) {
	arr := make([][]K, len(m))
	rm := make(map[K]uint16, len(m))
	i := 0
	for k, v := range m {
		arr[i] = make([]K, len(v)+1)
		arr[i][0] = k
		copy(arr[i][1:], v)
		rm[k] = uint16(i)
	}
	return ToSelfReferringArrUint16(arr, rm)
}

func ToSelfReferringArrUint16[A ~[][]K, M ~map[K]uint16, K comparable](a A, m M) ([][]uint16, map[K]uint16) {
	arr := make([][]uint16, len(a))
	for i := 0; i < len(a); i++ {
		arr[i] = make([]uint16, len(a[i])-1)
		for j := 0; j < len(arr[i]); j++ {
			arr[i][j] = m[a[i][j+1]]
		}
	}
	return arr, m
}

package util

import "fmt"

func ParseInt64(arr []byte) (res int64, err error) {
	var mult int64 = 1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] < '0' || arr[i] > '9' {
			err = fmt.Errorf("failed to parse %s to int", string(arr))
			return
		}
		res += int64(arr[i]-'0') * mult
		mult *= 10
	}
	return
}

func ParseInt64NoError(arr []byte) (res int64) {
	var mult int64 = 1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] < '0' || arr[i] > '9' {
			return
		}
		res += int64(arr[i]-'0') * mult
		mult *= 10
	}
	return
}

func ParseUint8NoError(arr []byte) (res uint8) {
	var mult uint8 = 1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] < '0' || arr[i] > '9' {
			return
		}
		res += (arr[i] - '0') * mult
		mult *= 10
	}
	return
}

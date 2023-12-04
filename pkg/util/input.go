package util

import (
	"bufio"
	"bytes"
	"io"
)

func DoEachRow(input []byte, f func(row []byte, nr int) error) (err error) {
	r := bufio.NewReader(bytes.NewReader(input))
	return DoEachRowReader(r, f)
}

func DoEachRowReader(r *bufio.Reader, f func(row []byte, nr int) error) (err error) {
	var (
		row  []byte
		done bool
	)
	for i := 0; !done; i++ {
		row, err = r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return
		}
		if err == io.EOF {
			done = true
		}
		err = f(row, i)
		if err != nil {
			return
		}
	}
	return
}

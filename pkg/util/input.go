package util

import (
	"bufio"
	"bytes"
	"io"
)

func DoEachRow(input []byte, f func(row []byte) error) (err error) {
	var (
		row  []byte
		done bool
	)
	r := bufio.NewReader(bytes.NewReader(input))
	for !done {
		row, err = r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return
		}
		if err == io.EOF {
			done = true
		}
		err = f(row)
		if err != nil {
			return
		}
	}
	return
}

func DoEachRowReader(r *bufio.Reader, f func(row []byte) error) (err error) {
	var (
		row  []byte
		done bool
	)
	for !done {
		row, err = r.ReadBytes('\n')
		if err != nil && err != io.EOF {
			return
		}
		if err == io.EOF {
			done = true
		}
		err = f(row)
		if err != nil {
			return
		}
	}
	return
}

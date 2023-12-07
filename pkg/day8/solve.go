package day8

import (
	"advent-of-code-2023/pkg/util"
	"embed"
)

//go:embed *.txt
var f embed.FS

func Solve() (err error) {
	var (
		b []byte
	)
	b, err = f.ReadFile("example.txt")
	if err != nil {
		return
	}

	err = util.DoEachRow(b, func(row []byte, nr int) error {

		return nil
	})
	if err != nil {
		return
	}

	return
}

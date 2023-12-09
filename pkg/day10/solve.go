package day10

import (
	"advent-of-code-2023/pkg/util"
	"embed"
	"fmt"
)

//go:embed *.txt
var f embed.FS

func Solve() (err error) {
	var (
		b   []byte
		row util.TokenSlice
	)
	b, err = f.ReadFile("example.txt")
	if err != nil {
		return
	}

	t := util.NewTokenParser(b)

	for t.More() {
		row = t.NextRow()
	}
	fmt.Println(row)
	return
}

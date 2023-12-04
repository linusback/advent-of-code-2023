package day4

import (
	"advent-of-code-2023/pkg/util"
	"embed"
	"fmt"
)

//go:embed *.txt
var f embed.FS

func Solve() (err error) {
	var (
		b       []byte
		total   uint64
		winning []uint8
		start   int
	)
	b, err = f.ReadFile("input.txt")
	if err != nil {
		return
	}
	for i := 0; i < len(b); i++ {
		if b[i] == ':' {
			winning = winning[:0]
			i++
			start = i
			for ; i < len(b); i++ {
				if b[i] == '|' {
					parseNumbers(&winning, b[start:i])
					i++
					break
				}
			}
			start = i
			for ; i < len(b); i++ {
				if b[i] == '\n' {
					total += checkWinningNumbers(b[start:i], winning)
					break
				}
			}

		}

	}
	fmt.Printf("total winnings %d\n", total)
	return
}

func checkWinningNumbers(src []byte, winning []uint8) uint64 {
	var won uint64
	for i := 0; i < len(src); i += 3 {
		if i+1 == len(src) && src[i] == ' ' {
			return won
		}
		if contains(winning, util.ParseUint8NoError(src[i:i+3])) {
			if won == 0 {
				won = 1
			} else {
				won = won << 1
			}

		}
	}
	return won
}

func parseNumbers(arr *[]uint8, src []byte) {
	for i := 0; i < len(src); i += 3 {
		if i+1 == len(src) && src[i] == ' ' {
			return
		}
		*arr = append(*arr, util.ParseUint8NoError(src[i:i+3]))
	}
}

func contains[K comparable](src []K, candidate K) bool {
	for i := 0; i < len(src); i++ {
		if src[i] == candidate {
			return true
		}
	}
	return false
}

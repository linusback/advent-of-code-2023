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
		b                 []byte
		total, totalCards uint64
		winning           []uint8
		cards             []uint64
		start             int
	)
	b, err = f.ReadFile("input.txt")
	if err != nil {
		return
	}
	cardIdx := -1
	for i := 0; i < len(b); i++ {
		if b[i] == ':' {
			cardIdx++
			if len(cards) <= cardIdx {
				cards = append(cards, 1)
			}
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
					total += checkWinningNumbers(b[start:i], winning, cardIdx, &cards)
					break
				}
			}

		}

	}
	//fmt.Printf("cards %v\n", cards)
	//fmt.Printf("cards %v\n", cards[:cardIdx+1])

	for i := 0; i <= cardIdx; i++ {
		totalCards += cards[i]
	}
	fmt.Printf("cards %v\n", totalCards)
	fmt.Printf("total winnings %d\n", total)
	return
}

func checkWinningNumbers(src []byte, winning []uint8, cardIdx int, cards *[]uint64) uint64 {
	var won uint64
	var matches int
	for i := 0; i < len(src); i += 3 {
		if i+1 == len(src) && src[i] == ' ' {
			addCards(matches, cardIdx, cards)
			return won
		}
		if contains(winning, util.ParseUint8NoError(src[i:i+3])) {
			matches++
			if won == 0 {
				won = 1
			} else {
				won = won << 1
			}

		}
	}
	addCards(matches, cardIdx, cards)
	return won
}

func addCards(matches, cardIdx int, cards *[]uint64) {
	multi := (*cards)[cardIdx]
	for i := 1; i <= matches; i++ {
		idx := cardIdx + i
		if len(*cards) <= idx {
			*cards = append(*cards, multi+1)
		} else {
			(*cards)[idx] += multi
		}
	}
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

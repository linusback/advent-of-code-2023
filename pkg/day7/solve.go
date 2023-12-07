package day7

import (
	"advent-of-code-2023/pkg/util"
	"embed"
	"fmt"
	"slices"
)

//go:embed *.txt
var f embed.FS

var order = []byte("23456789TJQKA")
var orderBytes [255]int

var order2 = []byte("J23456789TQKA")
var orderBytes2 [255]int

func populateOrder() {
	for i := 0; i < 13; i++ {
		orderBytes[order[i]] = i
		orderBytes2[order2[i]] = i
	}
}

type kind int

const (
	highCard kind = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type hand struct {
	kind, kind2 kind
	cards       [5]byte
	bid         uint64
}

func (h *hand) Print() {
	fmt.Printf("%s: %+v\n", string(h.cards[:]), h)
}

func Solve() (err error) {
	var (
		b                  []byte
		hands              = make([]hand, 0, 1000)
		h                  hand
		res, res2, handLen uint64
	)
	populateOrder()
	//b, err = f.ReadFile("example.txt")
	b, err = f.ReadFile("input.txt")
	if err != nil {
		return
	}

	err = util.DoEachRow(b, func(row []byte, nr int) error {
		h = hand{}
		copy(h.cards[:], row[:5])
		h.bid = util.ParseUint64IgnoreAll(row[5:])
		h.setKind()
		hands = append(hands, h)
		//h.Print()
		return nil
	})
	if err != nil {
		return
	}

	sortHands(hands)
	handLen = uint64(len(hands))
	for i := uint64(0); i < handLen; i++ {
		res += (i + 1) * hands[i].bid
	}

	sortHands2(hands)
	for i := uint64(0); i < handLen; i++ {
		res2 += (i + 1) * hands[i].bid
	}
	fmt.Println(res)
	fmt.Println(res2)

	return
}

var cards = [13]uint8{}
var cards2 = [13]uint8{}

func (h *hand) setKind() {
	clear(cards[:])
	clear(cards2[:])
	var jokers, first1, first2, second1, second2 uint8
	var firstB1, firstB2 byte
	for i := 0; i < len(h.cards); i++ {
		j := orderBytes[h.cards[i]]
		cards[j]++
		if h.cards[i] == 'J' {
			jokers++
		} else {
			cards2[j]++
		}
		if cards[j] > first1 {
			if h.cards[i] != firstB1 {
				second1 = first1
			}
			firstB1 = h.cards[i]
			first1 = cards[j]
		} else if cards[j] > second1 {
			second1 = cards[j]
		}

		if cards2[j] > first2 {
			if h.cards[i] != firstB2 {
				second2 = first2
			}
			firstB2 = h.cards[i]
			first2 = cards2[j]
		} else if cards2[j] > second2 {
			second2 = cards[j]
		}
	}

	h.kind = getKind(first1, second1, 0)
	h.kind2 = getKind(first2, second2, jokers)
}

func sortHands(hands []hand) {

	slices.SortFunc(hands, func(a, b hand) int {
		if a.kind != b.kind {
			return int(a.kind - b.kind)
		}
		for i := 0; i < 5; i++ {
			if a.cards[i] != b.cards[i] {
				return orderBytes[a.cards[i]] - orderBytes[b.cards[i]]
			}
		}
		return 0
	})
}
func sortHands2(hands []hand) {
	slices.SortFunc(hands, func(a, b hand) int {
		if a.kind2 != b.kind2 {
			return int(a.kind2 - b.kind2)
		}
		for i := 0; i < 5; i++ {
			if a.cards[i] != b.cards[i] {
				return orderBytes2[a.cards[i]] - orderBytes2[b.cards[i]]
			}
		}
		return 0
	})
}

func getKind(first, second, joker uint8) kind {
	switch first + joker {
	case 5:
		return fiveOfAKind
	case 4:
		return fourOfAKind
	case 3:
		if second == 2 {
			return fullHouse
		}
		return threeOfAKind
	case 2:
		if second == 2 {
			return twoPair
		}
		return onePair
	case 1:
		return highCard
	default:
		panic("invalid cards")
	}
}

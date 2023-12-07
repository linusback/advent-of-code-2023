package day7

import (
	"advent-of-code-2023/pkg/util"
	"bytes"
	"embed"
	"fmt"
	"slices"
)

//go:embed *.txt
var f embed.FS

var order = []byte("23456789TJQKA")
var order2 = []byte("J23456789TQKA")

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

func Solve() (err error) {
	var (
		b                  []byte
		hands              = make([]hand, 0, 1000)
		h                  hand
		res, res2, handLen uint64
	)
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
		return nil
	})
	if err != nil {
		return
	}
	//fmt.Println(hands)
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
	var jokers uint8
	for i := 0; i < len(h.cards); i++ {
		j := bytes.IndexByte(order, h.cards[i])
		cards[j]++
		if h.cards[i] == 'J' {
			jokers++
		} else {
			cards2[j]++
		}
	}

	slices.Sort(cards[:])
	slices.Sort(cards2[:])
	cards2[12] = cards2[12] + jokers
	h.kind = getKind(cards)
	h.kind2 = getKind(cards2)
}

func sortHands(hands []hand) {
	slices.SortFunc(hands, func(a, b hand) int {
		if a.kind != b.kind {
			return int(a.kind) - int(b.kind)
		}
		for i := 0; i < 5; i++ {
			if a.cards[i] == b.cards[i] {
				continue
			}
			ai := bytes.IndexByte(order, a.cards[i])
			bi := bytes.IndexByte(order, b.cards[i])
			return ai - bi
		}
		return 0
	})
}
func sortHands2(hands []hand) {
	slices.SortFunc(hands, func(a, b hand) int {
		if a.kind2 != b.kind2 {
			return int(a.kind2) - int(b.kind2)
		}
		for i := 0; i < 5; i++ {
			if a.cards[i] == b.cards[i] {
				continue
			}
			ai := bytes.IndexByte(order2, a.cards[i])
			bi := bytes.IndexByte(order2, b.cards[i])
			return ai - bi
		}
		return 0
	})
}

func getKind(cards [13]uint8) kind {
	switch cards[12] {
	case 5:
		return fiveOfAKind
	case 4:
		return fourOfAKind
	case 3:
		if cards[11] == 2 {
			return fullHouse
		}
		return threeOfAKind
	case 2:
		if cards[11] == 2 {
			return twoPair
		}
		return onePair
	case 1:
		return highCard
	default:
		panic("invalid cards")
	}
}

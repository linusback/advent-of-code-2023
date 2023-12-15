package day12

import (
	"advent-of-code-2023/pkg/util"
	"embed"
	"fmt"
)

//go:embed *.txt
var f embed.FS

type springRow struct {
	springs      []byte
	springGroups []uint64
	groups       []group
}

func (r springRow) findWorkingCombinations() (res uint64) {
	fmt.Println(string(r.springs))
	fmt.Println(r.springGroups)
	return 0
}

type group struct {
	raw       []byte
	perm      [][]byte
	wildCards []int
	offset    int
}

func Solve() (res1t, res2t uint64, err error) {
	var (
		b   []byte
		row util.TokenSlice
	)
	b, err = f.ReadFile("example.txt")
	if err != nil {
		return
	}

	t := util.NewTokenParserSeparators(b, ' ', ',')
	springs := make([]springRow, 0, 1000)
	for t.More() {
		row = t.NextRow()
		s := springRow{
			springs: row[0],
		}
		//if len(springs) == 0 {
		//
		//}
		//combinations(row[0])
		s.springs = row[0]
		s.groups = combinations(row[0])
		for i := 1; i < len(row); i++ {
			s.springGroups = append(s.springGroups, row[i].ParseUInt64())
		}
		s.findWorkingCombinations()
		springs = append(springs, s)
	}

	//fmt.Println(springs)
	return
}

func combinations(arr []byte) []group {
	res := make([]group, 0, len(arr)/2)
	wildCards := make([]int, 0, len(arr))
	insideGroup := false
	s := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == '.' {
			if insideGroup {
				insideGroup = false
				curr := group{
					raw:       arr[s:i],
					offset:    s,
					wildCards: make([]int, len(wildCards)),
				}
				curr.setWildCards(wildCards)
				curr.generatePermutations()
				res = append(res, curr)
				wildCards = wildCards[:0]
			}

			continue
		}
		if !insideGroup {
			s = i
			insideGroup = true
		}
		if arr[i] == '?' {
			wildCards = append(wildCards, i)
		}

	}
	if insideGroup {
		curr := group{
			raw:       arr[s:],
			offset:    s,
			wildCards: make([]int, len(wildCards)),
		}
		curr.setWildCards(wildCards)
		curr.generatePermutations()
		res = append(res, curr)
	}
	return res

}

func (g *group) generatePermutations() {
	if len(g.wildCards) == 0 {
		g.perm = [][]byte{g.raw}
		return
	}

	pLen := 1 << len(g.wildCards)
	g.perm = make([][]byte, 0, pLen)
	start := make([]byte, len(g.wildCards))
	for i := 0; i < len(start); i++ {
		start[i] = '.'
	}
	g.perm = append(g.perm, start)
	g.perm = wildcardPerm(g.perm, start, 0, len(g.wildCards))
}

func wildcardPerm(perm [][]byte, base []byte, i, len int) [][]byte {
	if i == len {
		return perm
	}
	next1 := make([]byte, len)
	copy(next1, base)
	next1[i] = '#'
	perm = append(perm, next1)
	perm = wildcardPerm(perm, base, i+1, len)
	perm = wildcardPerm(perm, next1, i+1, len)
	return perm
}

func (g *group) setWildCards(w []int) {
	for i := 0; i < len(g.wildCards); i++ {
		g.wildCards[i] = w[i] - g.offset
	}
}

// ???
// 111
// 222

// 112
// 211
// 121

// 221
// 122
// 212

//

// ????

// 1111
// replace one
// 2111
// 1211
// 1121
// 1112

// replace one
// 2211
// 1221
// 1122
// 2112

// 2222

// 112
// 211
// 121

// 221
// 122
// 212

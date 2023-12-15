package day12

import (
	"advent-of-code-2023/pkg/util"
	"embed"
	"slices"
)

//go:embed *.txt
var f embed.FS

type springRow struct {
	springs      []byte
	springGroups []uint64
	groups       []group
	line         int
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
	b, err = f.ReadFile("input.txt")
	if err != nil {
		return
	}

	//permu := util.Permutate([][]int{{1, 2}, {1, 2}, {1, 2}})
	//fmt.Println(permu)
	//return
	t := util.NewTokenParserSeparators(b, ' ', ',')
	springs := make([]springRow, 0, 1000)
	line := 0
	for t.More() {
		row = t.NextRow()
		s := springRow{
			springs: row[0],
			line:    line,
		}
		line++
		//if len(springs) == 0 {
		//
		//}
		//combinations(row[0])
		s.springs = row[0]
		s.groups = combinations(row[0])
		for i := 1; i < len(row); i++ {
			s.springGroups = append(s.springGroups, row[i].ParseUInt64())
		}
		c := s.findWorkingCombinations()
		res1t += c
		//fmt.Println(string(row[0]), c)
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
	g.perm = g.expandWildcard()
}

func (g *group) expandWildcard() [][]byte {
	for i := 0; i < len(g.perm); i++ {
		r := make([]byte, len(g.raw))
		copy(r, g.raw)
		for j := 0; j < len(g.perm[i]); j++ {
			b := g.perm[i][j]
			idx := g.wildCards[j]
			r[idx] = b
		}
		g.perm[i] = r
	}
	return g.perm
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

func (g *group) getPossibleSpringGroups() (res [][]uint64) {
	res = make([][]uint64, 0, len(g.perm))
	for i := 0; i < len(g.perm); i++ {
		r := numberOfHashtag(g.perm[i])
		if len(r) > 0 {
			res = append(res, r)
		}
	}
	return res
}

func numberOfHashtag(arr []byte) (res []uint64) {
	inside := false
	s := 0
	//fmt.Println("checking hashtags in: ", string(arr))
	for i := 0; i < len(arr); i++ {
		if inside && arr[i] == '.' {
			res = append(res, uint64(i-s))
			inside = false
			continue
		}
		if !inside && arr[i] == '#' {
			inside = true
			s = i
		}
	}
	if inside {
		res = append(res, uint64(len(arr)-s))
	}
	//fmt.Println("hashtags, ", res)
	return
}

func (r *springRow) findWorkingCombinations() (res uint64) {
	arr := make([][][]uint64, 0, len(r.groups))
	resLen := uint64(1)
	for i := 0; i < len(r.groups); i++ {
		g := r.groups[i].getPossibleSpringGroups()
		resLen *= uint64(len(g))
		arr = append(arr, g)
	}
	candidates := make([][]uint64, 0, resLen)
	candidates = generateCombinations(candidates, arr, 0)
	return validCombinations(candidates, r.springGroups)
}

func validCombinations(candidates [][]uint64, groups []uint64) (valid uint64) {
	for i := 0; i < len(candidates); i++ {
		if slices.Equal(candidates[i], groups) {
			valid++
		}
	}
	return
}

func generateCombinations(res [][]uint64, arr [][][]uint64, i int) [][]uint64 {
	if i == len(arr) {
		return res
	}
	if i == 0 {
		for j := 0; j < len(arr[0]); j++ {
			res = append(res, arr[0][j])
		}
		return generateCombinations(res, arr, i+1)
	}
	res2 := make([][]uint64, 0, len(res)*len(arr[i]))
	for k := 0; k < len(res); k++ {
		for j := 0; j < len(arr[i]); j++ {
			r := append(res[k], arr[i][j]...)
			res2 = append(res2, r)
		}
	}
	return generateCombinations(res2, arr, i+1)
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

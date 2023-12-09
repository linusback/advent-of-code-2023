package day8Alt

import (
	"advent-of-code-2023/pkg/util"
	"embed"
	"fmt"
	"time"
)

//go:embed *.txt
var f embed.FS

type network struct {
	instructions []uint16
	nodes        [758][2]uint16
}

// Path describes a path, visited and cycle can be removed assuming only 1 exit node per path
type Path struct {
	start, curr uint16
	toFind      uint64
}

func Solve() (err error) {
	var (
		b     []byte
		r     util.TokenSlice
		n     network
		node  uint16
		paths []Path
	)

	b, err = f.ReadFile("input.txt")
	//b, err = f.ReadFile("example.txt")
	if err != nil {
		return
	}

	p := util.NewTokenParserWithSeparators(b, util.Newline, ' ', '=', '(', ')', ',')

	// left right parse
	r = p.NextRow()
	n.instructions = make([]uint16, len(r[0]))
	for i := 0; i < len(r[0]); i++ {
		switch r[0][i] {
		case 'L':
			n.instructions[i] = 0
		case 'R':
			n.instructions[i] = 1
		default:
			panic(fmt.Sprintf("%c is not a left right instruction", r[0][i]))

		}
	}
	nMap := make(map[string]uint16, 758)
	arr := make([][]string, 758)
	paths = make([]Path, 0, 6)
	exitNodes := uint16(0)
	z := uint16(6)
	for p.More() {
		r = p.NextRow()
		if r[0][2] == 'A' {
			node = uint16(len(paths))
			paths = append(paths, Path{
				start: node,
				curr:  node,
			})
			arr[node] = []string{string(r[0]), string(r[1]), string(r[2])}
			nMap[string(r[0])] = node
			continue
		}

		if r[0][2] == 'Z' {
			node = uint16(len(arr)) - 1 - exitNodes
			arr[node] = []string{string(r[0]), string(r[1]), string(r[2])}
			nMap[string(r[0])] = node
			exitNodes++
			continue
		}
		node = z
		arr[node] = []string{string(r[0]), string(r[1]), string(r[2])}
		nMap[string(r[0])] = node
		z++

	}
	var arr2 [][]uint16
	arr2, nMap = util.ToSelfReferringArrUint16(arr, nMap)
	for i := 0; i < len(arr2); i++ {
		copy(n.nodes[i][:], arr2[i])
	}

	//var left, right string
	//for s, u := range nMap {
	//	if s[2] == 'Z' {
	//
	//		for s2, u2 := range nMap {
	//			if u2 == n.nodes[u][0] {
	//				left = s2
	//			}
	//			if u2 == n.nodes[u][1] {
	//				right = s2
	//			}
	//		}
	//		fmt.Println(s, ": ", u, " = ", left, ", ", right)
	//	}
	//}
	//fmt.Println(len(arr) - 1 - int(exitNodes))
	//return

	start := time.Now()

	//part 1
	var total1 uint64
	next := nMap["AAA"]
	exitNode := nMap["ZZZ"]
	for i := 0; next != exitNode; i++ {
		if i == len(n.instructions) {
			i = 0
		}
		total1++
		next = n.nodes[next][n.instructions[i]]
	}
	fmt.Printf("part1 in %v: %d\n", time.Since(start), total1)
	start = time.Now()

	// part2
	var total2 uint64
	var pa *Path
	var found int
	normalNodes := uint16(len(arr)) - 1 - exitNodes
	done := make([]Path, 0, len(paths))
	start2 := time.Now()
	for i := 0; total2 <= 14449445933179; i++ {
		if i == len(n.instructions) {
			i = 0
		}
		total2++
		found = 0
		if total2%100000000 == 0 {
			fmt.Printf("at %d it took %v\n", total2, time.Since(start2))
			start2 = time.Now()
		}
		for j := 0; j < len(paths); j++ {
			pa = &paths[j]
			pa.curr = n.nodes[pa.curr][n.instructions[i]]
			if pa.curr > normalNodes {
				found++
				//pa.toFind = total2
				//done = append(done, *pa)
				//if j+1 < len(paths) {
				//	// move one instead of all
				//	paths[j] = paths[len(paths)-1]
				//}
				//paths = paths[:len(paths)-1]
				//j--
			}
		}
		if found == len(paths) {
			break
		}
	}
	var toFindMul = uint64(len(n.instructions))
	for _, p1 := range done {
		toFindMul = util.Lcd(toFindMul, p1.toFind)
	}
	// this might find wrong answer since it could be earlier if a long cycle contains more than one exit node
	fmt.Printf("part2 in %v: %d\n", time.Since(start), toFindMul)

	return
}

func printStart(paths []Path) {
	s := make([]string, 0, len(paths))
	for _, p := range paths {
		s = append(s, fmt.Sprintf("%d", p.start))
	}
	fmt.Println(s)
}

const (
	hundreds = uint16(26 * 26)
	tens     = uint16(26)
)

// change to uint16 with base 26 so AAA = 0 AND ZZZ = 17575 (26*26*26 - 1)
func toUint16(b []byte) uint16 {
	var r uint16
	if len(b) != 3 {
		panic(fmt.Sprintf("wrong length of %s", string(b)))
	}
	r += uint16(b[0]-'A') * hundreds
	r += uint16(b[1]-'A') * tens
	r += uint16(b[2] - 'A')
	return r
}

func toString(u uint16) string {
	var b [3]byte
	b[2] = byte(u%tens) + 'A'
	b[1] = byte((u%hundreds)/tens) + 'A'
	b[0] = byte(u/hundreds) + 'A'
	return string(b[:])
}

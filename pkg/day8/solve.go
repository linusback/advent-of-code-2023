package day8

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
	nodes        [17576][2]uint16
}

// Path describes a path, visited and cycle can be removed assuming only 1 exit node per path
type Path struct {
	start, curr, end uint16
	visited          [17576]uint64
	cycle, toFind    uint64
}

func (p *Path) ToString() string {

	visited := make([]string, p.cycle)
	visited[0] = toString(p.start)
	found := 1
	for i := uint16(0); i < uint16(len(p.visited)); i++ {
		if p.visited[i] > 0 {
			visited[p.visited[i]] = toString(i)
			found++
		}
	}
	visited2 := make([]string, 0, found)
	for i := 0; i < len(visited); i++ {
		if visited[i] != "" {
			visited2 = append(visited2, visited[i])
		}
	}
	return fmt.Sprintf("{start: %s, curr: %s, end: %s, toFind:%d, cycle: %d, visited: %v}", toString(p.start), toString(p.curr), toString(p.end), p.toFind, p.cycle, visited2)
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
	paths = make([]Path, 0, 6)
	for p.More() {
		r = p.NextRow()
		node = toUint16(r[0])
		n.nodes[node][0] = toUint16(r[1])
		n.nodes[node][1] = toUint16(r[2])
		if node%tens == 0 {
			paths = append(paths, Path{
				start: node,
				curr:  node,
			})
		}
	}
	start := time.Now()
	//part 1
	var total1 uint64
	next := uint16(0)
	for i := 0; next < 17575; i++ {
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
	done := make([]Path, 0, len(paths))
	for i := 0; 0 < len(paths); i++ {
		if i == len(n.instructions) {
			i = 0
		}
		total2++
		for j := 0; j < len(paths); j++ {
			pa = &paths[j]
			pa.curr = n.nodes[pa.curr][n.instructions[i]]
			if pa.curr%tens == 25 && pa.end == 0 {
				pa.toFind = total2
				pa.end = pa.curr
			}
			if pa.end > 0 && pa.visited[pa.curr] > 0 {
				pa.cycle = total2
				done = append(done, *pa)
				if j+1 < len(paths) {
					copy(paths[j:], paths[j+1:])
				}
				paths = paths[:len(paths)-1]
				j--

			} else {
				pa.visited[pa.curr] = total2
			}
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

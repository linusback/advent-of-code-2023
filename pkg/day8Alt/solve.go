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
	start, curr          uint16
	cycleEnds, cycleLens []uint64
	cycles               []string
	cycle                []uint16
}

func Solve() (err error) {
	var (
		b     []byte
		r     util.TokenSlice
		n     network
		node  uint16
		paths []Path
		arr2  [][]uint16
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
		switch r[0][2] {
		case 'A':
			node = uint16(len(paths))
			paths = append(paths, Path{
				start: node,
				curr:  node,
			})
		case 'Z':
			node = uint16(len(arr)) - 1 - exitNodes
			exitNodes++
		default:
			node = z
			z++
		}
		arr[node] = []string{string(r[0]), string(r[1]), string(r[2])}
		nMap[string(r[0])] = node

	}

	arr2, nMap = util.ToSelfReferringArrUint16(arr, nMap)
	for i := 0; i < len(arr2); i++ {
		copy(n.nodes[i][:], arr2[i])
	}

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
	var (
		pa                         *Path
		cycStr                     string
		total2, cycleEnd, cycleLen uint64
	)
	normalNodes := uint16(len(arr)) - 1 - exitNodes
	done := make([]Path, 0, len(paths))
	for i := 0; 0 < len(paths); i++ {
		if i == len(n.instructions) {
			i = 0
		}
		total2++
	pathLoop:
		for j := 0; j < len(paths); j++ {
			pa = &paths[j]
			pa.cycle = append(pa.cycle, pa.curr)
			pa.curr = n.nodes[pa.curr][n.instructions[i]]
			if pa.curr > normalNodes {
				cycleEnd = uint64(i)
				cycStr = util.Uint16ToString(pa.cycle)
				cycleLen = uint64(len(pa.cycle))

				pa.cycle = pa.cycle[:0]
				for k := 0; k < len(pa.cycleEnds); k++ {
					if pa.cycleEnds[k] == cycleEnd && pa.cycles[k] == cycStr {
						done = append(done, *pa)
						if j+1 < len(paths) {
							// move one instead of all
							paths[j] = paths[len(paths)-1]
						}
						paths = paths[:len(paths)-1]
						j--
						continue pathLoop
					}
				}

				pa.cycleEnds = append(pa.cycleEnds, uint64(i))
				pa.cycles = append(pa.cycles, cycStr)
				pa.cycleLens = append(pa.cycleLens, cycleLen)
			}
		}
	}

	numbers := make([][]uint64, len(done))
	var (
		toFindMul, result2 uint64
		cycles             []uint64
		last               uint64
	)
	for i := 0; i < len(done); i++ {
		cycles = util.FilterUnique(done[i].cycleLens)
		last = 0
		for j := 0; j < len(cycles); j++ {
			cycles[j] += last
			last = cycles[j]

		}
		numbers[i] = cycles
	}
	numbers = util.Permutate(numbers)
	//fmt.Println(numbers)
	result2 = ^uint64(0)
	for i := 0; i < len(numbers); i++ {
		toFindMul = uint64(len(n.instructions))
		for j := 0; j < len(numbers[i]); j++ {
			toFindMul = util.Lcd(toFindMul, numbers[i][j])
		}
		if result2 > toFindMul {
			result2 = toFindMul
		}
	}
	// this might find wrong answer since it could be earlier if a long cycle contains more than one exit node
	fmt.Printf("part2 in %v: %d\n", time.Since(start), toFindMul)

	return
}

func ToString(u []uint16, m *map[uint16]string) string {
	s := make([]byte, len(u)*3)
	for i := 0; i < len(u); i++ {
		copy(s[i*3:], (*m)[u[i]])
	}
	return string(s)
}

func printStart(paths []Path) {
	s := make([]string, 0, len(paths))
	for _, p := range paths {
		s = append(s, fmt.Sprintf("%d", p.start))
	}
	fmt.Println(s)
}

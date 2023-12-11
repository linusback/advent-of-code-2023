package day10

import (
	"advent-of-code-2023/pkg/util"
	"embed"
	"fmt"
	"slices"
	"time"
)

//go:embed *.txt
var f embed.FS

type TileType uint8

const (
	unknown TileType = iota
	pipe
	outside
	inside
)

type Tile struct {
	x, y      int64
	b         byte
	connected [][2]int64
	val       uint64
	tt        TileType

	top, bottom, right, left int
}

func (ts *Tiles) GetCorr(x, y int) *Tile {
	return &((*ts)[y][x])
}

func (ts *Tiles) Get(corr [2]int64) *Tile {
	return &((*ts)[corr[1]][corr[0]])
}

type Tiles [][]Tile

const maxUint64 = ^uint64(0)
const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

func Solve() (err error) {
	var (
		b     []byte
		tiles Tiles
		start *Tile
	)
	startTime := time.Now()
	b, err = f.ReadFile("input.txt")
	if err != nil {
		return
	}

	err = util.DoEachRowAll(b, func(row []byte, rows [][]byte, nr, total int) error {
		if nr == 0 {
			tiles = make(Tiles, total)
		}
		tiles[nr] = make([]Tile, len(row))
		for x := 0; x < len(row); x++ {
			tiles[nr][x].b = row[x]
			tiles[nr][x].x = int64(x)
			tiles[nr][x].y = int64(nr)
			tiles[nr][x].y = int64(nr)
			tiles[nr][x].val = maxUint64
			if row[x] == 'S' {
				start = &tiles[nr][x]
				start.val = 0
				start.tt = pipe
			}
			tiles[nr][x].CreateConnected(int64(len(row)), int64(total))
		}
		return nil
	})
	if err != nil {
		return
	}
	start.SetConnectedForStart(&tiles)
	//fmt.Println(start)
	start.SetPipeStructStart()
	//fmt.Println(start)
	fmt.Println("parsing: ", time.Since(startTime))
	startTime = time.Now()
	//part 1
	result1 := part1(&tiles, start)
	fmt.Printf("part1: %d in: %v\n", result1, time.Since(startTime))
	startTime = time.Now()

	//tiles.PrintVals()
	//fmt.Println()
	//tiles.PrintTypes()
	result2 := tiles.SetTypes()
	//tiles.PrintTypes()
	fmt.Printf("part2: %d in: %v\n", result2, time.Since(startTime))

	return
}

func part1(t *Tiles, curr *Tile) uint64 {
	var maxVal uint64
	next := curr.GetConnected(nil, t)
	for len(next) > 0 {
		curr = next[0]
		curr.tt = pipe
		curr.SetPipeStruct()
		next = slices.Delete(next, 0, 1)
		if curr.val > maxVal {
			maxVal = curr.val
		}
		next = append(next, curr.GetConnected(curr, t)...)
	}
	return maxVal
}

func (ts *Tiles) SetTypes() uint64 {
	var (
		top, bottom, minH, minV int
		t                       *Tile
		result                  = uint64(0)
	)
	columns := make([][2]int, len((*ts)[0]))
	for y := 0; y < len(*ts); y++ {
		top, bottom = 0, 0
		//fmt.Println("y: ", y, len((*ts)[y])-1)
		//fmt.Println("x|top,bot,lef,rig")
		for x := 0; x < len((*ts)[y]); x++ {
			t = ts.GetCorr(x, y)
			if t.tt == unknown {
				t.tt = outside
				minH, minV = columns[x][0], top
				if minV > bottom {
					minV = bottom
				}
				if minH > columns[x][1] {
					minH = columns[x][1]
				}
				//if y == 6 {
				//
				//	fmt.Printf("%d not|  %d,  %d,  %d,  %d\n", t.x, top, bottom, columns[x][0], columns[x][1])
				//}
				//fmt.Println(y, x, minV, minH)
				if minV%2 == 1 && minH%2 == 1 {
					t.tt = inside
					result++
				}
				continue
			}
			//fmt.Printf("%d|  %d,  %d,  %d,  %d,  %c\n", t.x, t.top, t.bottom, t.left, t.right, t.b)
			top += t.top
			bottom += t.bottom
			columns[x][0] += t.left
			columns[x][1] += t.right
		}
	}
	//fmt.Println(columns)
	return result
}

func (t *Tile) SetPipeStruct() {
	switch t.b {
	case '|':
		t.top = 1
		t.bottom = 1
	case '-':
		t.left = 1
		t.right = 1
	case 'L':
		t.top = 1
		t.right = 1
	case 'J':
		t.top = 1
		t.left = 1
	case '7':
		t.bottom = 1
		t.left = 1
	case 'F':
		t.bottom = 1
		t.right = 1
	}
}

func (t *Tile) SetPipeStructStart() {
	possible := [4][2]int64{
		{0, -1},
		{1, 0},
		{-1, 0},
		{0, 1},
	}
	t.top = 1
	t.bottom = 1
	t.right = 1
	t.left = 1

	for i := 0; i < len(possible); i++ {
		if !slices.Contains(t.connected, [2]int64{t.x + possible[i][0], t.y + possible[i][1]}) {
			switch i {
			case 0:
				t.top = 0
			case 1:
				t.right = 0
			case 2:
				t.left = 0
			case 3:
				t.bottom = 0
			}
		}
	}
}

func (t *Tile) CreateConnected(lenX, lenY int64) {
	t.connected = make([][2]int64, 2)
	switch t.b {
	case '|':
		t.connected[0][0] = t.x
		t.connected[0][1] = t.y - 1
		t.connected[1][0] = t.x
		t.connected[1][1] = t.y + 1
	case '-':
		t.connected[0][0] = t.x - 1
		t.connected[0][1] = t.y
		t.connected[1][0] = t.x + 1
		t.connected[1][1] = t.y
	case 'L':
		t.connected[0][0] = t.x
		t.connected[0][1] = t.y - 1
		t.connected[1][0] = t.x + 1
		t.connected[1][1] = t.y
	case 'J':
		t.connected[0][0] = t.x
		t.connected[0][1] = t.y - 1
		t.connected[1][0] = t.x - 1
		t.connected[1][1] = t.y
	case '7':
		t.connected[0][0] = t.x
		t.connected[0][1] = t.y + 1
		t.connected[1][0] = t.x - 1
		t.connected[1][1] = t.y
	case 'F':
		t.connected[0][0] = t.x
		t.connected[0][1] = t.y + 1
		t.connected[1][0] = t.x + 1
		t.connected[1][1] = t.y
	case '.':
		t.connected = t.connected[:0]
	case 'S':
		t.connected = make([][2]int64, 8)
		t.connected[0][1] = t.y - 1
		t.connected[0][0] = t.x

		t.connected[1][1] = t.y
		t.connected[1][0] = t.x - 1

		t.connected[2][1] = t.y
		t.connected[2][0] = t.x + 1

		t.connected[3][1] = t.y + 1
		t.connected[3][0] = t.x
	default:
		panic(fmt.Sprintf("no support for %c", t.b))
	}
	t.connected = slices.DeleteFunc(t.connected, func(corr [2]int64) bool {
		return corr[0] < 0 || corr[1] < 0 || corr[0] >= lenX || corr[1] >= lenY
	})
}

func (t *Tile) SetConnectedForStart(ts *Tiles) {
	var (
		toDelete [][2]int64
		t2       *Tile
		corr     = [2]int64{t.x, t.y}
	)
	for i := 0; i < len(t.connected); i++ {
		t2 = ts.Get(t.connected[i])
		if !slices.Contains(t2.connected, corr) {
			toDelete = append(toDelete, t.connected[i])
		}
	}
	t.connected = slices.DeleteFunc(t.connected, func(curr [2]int64) bool {
		return slices.Contains(toDelete, curr)
	})
}

func (t *Tile) GetConnected(curr *Tile, ts *Tiles) []*Tile {
	res := make([]*Tile, 0, len(t.connected))
	var candidate *Tile
	for i := 0; i < len(t.connected); i++ {
		if curr != nil && t.connected[i][0] == curr.x && t.connected[i][1] == curr.y {
			continue
		}
		candidate = ts.Get(t.connected[i])
		if curr != nil && candidate.val <= t.val+1 {
			continue
		}
		candidate.val = t.val + 1
		res = append(res, candidate)
	}
	return res
}

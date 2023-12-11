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
}

type Tiles [][]Tile

const maxUint64 = ^uint64(0)

func Solve() (err error) {
	var (
		b     []byte
		tiles Tiles
		start *Tile
	)
	startTime := time.Now()
	b, err = f.ReadFile("example22.txt")
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
	fmt.Println("parsing: ", time.Since(startTime))
	startTime = time.Now()
	//part 1
	result1 := part1(&tiles, start)
	fmt.Printf("part1: %d in: %v\n", result1, time.Since(startTime))
	startTime = time.Now()

	tiles.PrintVals()
	fmt.Println()
	tiles.PrintTypes()
	result2 := 1
	fmt.Printf("part2: %d in: %v\n", result2, time.Since(startTime))

	return
}

func part1(t *Tiles, curr *Tile) uint64 {
	var maxVal uint64
	next := curr.GetConnected(nil, t)
	for len(next) > 0 {
		curr = next[0]
		curr.tt = pipe
		next = slices.Delete(next, 0, 1)
		if curr.val > maxVal {
			maxVal = curr.val
		}
		next = append(next, curr.GetConnected(curr, t)...)
	}
	return maxVal
}

func (ts *Tiles) PrintVals() {
	row := make([]byte, 2*len((*ts)[0])+1)
	for i := 0; i < len(*ts); i++ {
		row = row[:0]
		row = append(row, '[')
		for j := 0; j < len((*ts)[i]); j++ {
			if (*ts)[i][j].val < maxUint64 {
				row = append(row, []byte(fmt.Sprintf("%d", (*ts)[i][j].val))...)
			} else {
				row = append(row, ' ')
			}
			row = append(row, ',')
		}
		row[len(row)-1] = ']'
		fmt.Println(string(row))
	}
}

func (ts *Tiles) PrintTypes() {
	var b byte
	var s string
	row := make([]byte, 2*len((*ts)[0])+1)
	for i := 0; i < len(*ts); i++ {
		row = row[:0]
		for j := 0; j < len((*ts)[i]); j++ {
			switch (*ts)[i][j].tt {
			case unknown:
				b = 'U'
			case pipe:
				switch (*ts)[i][j].b {
				case '|':
					s = "┃ "
				case '-':
					s = "━━"
				case 'L':
					s = "┗━"
				case 'J':
					s = "┛ "
				case '7':
					s = "┓ "
				case 'F':
					s = "┏━"
				case 'S':
					s = handleStart(&(*ts)[i][j])
				}
				row = append(row, []byte(s)...)
				continue
			case outside:
				b = 'O'
			case inside:
				b = 'I'
			}
			row = append(row, b, ' ')
		}
		fmt.Println(string(row))
	}
}

func handleStart(start *Tile) string {
	switch len(start.connected) {
	case 4:
		return "╋━"
	case 3:
		return missing(start.connected, [2]int64{start.x, start.y})
	case 2:
		return missing2(start.connected, [2]int64{start.x, start.y})
	default:
		return "S "
	}
}

func missing(conn [][2]int64, self [2]int64) string {
	var missingVal int
	possible := [4][2]int64{
		{0, -1},
		{1, 0},
		{-1, 0},
		{0, 1},
	}
	for i := 0; i < len(possible); i++ {
		if !slices.Contains(conn, [2]int64{self[0] + possible[i][0], self[1] + possible[i][1]}) {
			missingVal = i
			break
		}
	}
	switch missingVal {
	case 0:
		return "┳━"
	case 1:
		return "┫ "
	case 2:
		return "┣━"
	case 3:
		return "┻━"
	default:
		return "S "
	}
}

func missing2(conn [][2]int64, self [2]int64) string {
	var missingVal1 = -1
	var missingVal2 = -1
	possible := [4][2]int64{
		{0, -1},
		{1, 0},
		{-1, 0},
		{0, 1},
	}
	for i := 0; i < len(possible); i++ {
		if missingVal1 == -1 && !slices.Contains(conn, [2]int64{self[0] + possible[i][0], self[1] + possible[i][1]}) {
			missingVal1 = i
			continue
		}
		if missingVal2 == -1 && !slices.Contains(conn, [2]int64{self[0] + possible[i][0], self[1] + possible[i][1]}) {
			missingVal2 = i
			break
		}
	}
	if missingVal1 == 1 && missingVal2 == 2 {
		return "┃ "
	}
	if missingVal1 == 0 && missingVal2 == 3 {
		return "━━"
	}
	if missingVal1 == 2 && missingVal2 == 3 {
		return "┗━"
	}
	if missingVal1 == 1 && missingVal2 == 3 {
		return "┛ "
	}
	if missingVal1 == 0 && missingVal2 == 1 {
		return "┓ "
	}
	if missingVal1 == 0 && missingVal2 == 2 {
		return "┏━"
	}
	return "S "
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

func (ts *Tiles) Get(corr [2]int64) *Tile {
	return &((*ts)[corr[1]][corr[0]])
}

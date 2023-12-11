package day10

import (
	"fmt"
	"slices"
)

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

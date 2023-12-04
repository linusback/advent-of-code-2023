package day3

import (
	"advent-of-code-2023/pkg/util"
	"embed"
	"fmt"
)

//go:embed *.txt
var f embed.FS

type Gear struct {
	X, Y    int
	numbers []int64
}

func Solve() (err error) {
	var (
		b                    []byte
		number, sum, product int64
		gears                []Gear
	)
	b, err = f.ReadFile("input.txt")

	start := -1
	err = util.DoEachRowAll(b, func(row []byte, all [][]byte, nr, total int) (err2 error) {
		start = -1
		for i := 0; i < len(row); i++ {
			if '0' <= row[i] && row[i] <= '9' && start == -1 {
				start = i
			}
			if start > -1 {
				if row[i] < '0' || row[i] > '9' {
					number, err2 = util.ParseInt64(row[start:i])
				} else if i+1 == len(row) {
					number, err2 = util.ParseInt64(row[start : i+1])
				} else {
					continue
				}
				if err2 != nil {
					return
				}
				if ValidNumber(start, i, len(row), nr, total, all, number, &gears) {
					sum += number
				}
				start = -1
			}

		}

		return
	})
	for i := 0; i < len(gears); i++ {
		if len(gears[i].numbers) == 2 {
			product += gears[i].numbers[0] * gears[i].numbers[1]
		}
	}
	fmt.Println("product is", product)
	fmt.Printf("sum is %d\n", sum)
	return
}

func ValidNumber(start, end, colLen, row, rowLen int, rows [][]byte, number int64, gears *[]Gear) bool {
	var res bool
	if start > 0 {
		start -= 1
	}
	if end < colLen {
		end += 1
	}
	res = IsSymbol(rows[row][start:start+1], start, row, number, gears) || res
	res = IsSymbol(rows[row][end-1:end], end-1, row, number, gears) || res

	if row > 0 {
		res = IsSymbol(rows[row-1][start:end], start, row-1, number, gears) || res
	}
	if row < rowLen-1 {
		res = IsSymbol(rows[row+1][start:end], start, row+1, number, gears) || res
	}
	return res

}

func IsSymbol(src []byte, x, y int, number int64, gears *[]Gear) bool {
	res := false
	for i := 0; i < len(src); i++ {
		switch src[i] {
		case '.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case '*':
			res = true
			idx := findGear(gears, x+i, y)
			if idx == -1 {
				*gears = append(*gears, Gear{
					X:       x + i,
					Y:       y,
					numbers: []int64{number},
				})
			} else {
				(*gears)[idx].numbers = append((*gears)[idx].numbers, number)
			}
		default:
			res = true
		}
	}
	return res
}

func findGear(gears *[]Gear, x int, y int) int {
	for i := 0; i < len(*gears); i++ {
		if (*gears)[i].X == x && (*gears)[i].Y == y {
			return i
		}
	}
	return -1
}

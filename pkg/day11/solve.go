package day11

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"strconv"
)

//go:embed *.txt
var f embed.FS

type point struct {
	x, y, name int
}

var (
	columnMap, rowMap [140][140]int
	columns, rows     [140]int
)

const galaxyByte byte = '#'

func Solve() (res1, res2 int64, err error) {
	var (
		fil fs.File
		gal []point
	)
	fil, err = f.Open("input.txt")
	if err != nil {
		return
	}
	tiles := [140][140]byte{}

	//fil, err = f.Open("example.txt")
	//if err != nil {
	//	return
	//}
	//tiles := [][]byte{}

	//start := time.Now()
	buff := make([]byte, 141)
	n := 0
	for i := 0; i < len(tiles); i++ {
		n, err = fil.Read(buff)
		if n < 140 {
			err = fmt.Errorf("failed to read enougth row: %d, n: %d", i, n)
		}
		if err != nil && err != io.EOF {
			return
		}
		copy(tiles[i][:], buff)
	}
	//tiles := bytes.Split(b, []byte{'\n'})
	galaxies := 0
	//fmt.Printf("setup %v\n", time.Since(start))
	//start = time.Now()
	gal = make([]point, 0, 440)
	for y := 0; y < len(tiles); y++ {
		galaxies = 0
		for x := 0; x < len(tiles[y]); x++ {
			if tiles[y][x] == galaxyByte {
				galaxies++
				columns[x]++
				gal = append(gal, point{
					x:    x,
					y:    y,
					name: len(gal) + 1,
				})
			}
		}
		rows[y] = galaxies
	}
	for x := 0; x < len(columns); x++ {
		if columns[x] > 0 {
			columns[x] = 0
			continue
		}
		columns[x] = 1
	}
	for y := 0; y < len(rows); y++ {
		if rows[y] > 0 {
			rows[y] = 0
			continue
		}
		rows[y] = 1
	}

	//columnMap := createPairs(columns)
	//rowMap := createPairs(rows)

	createPairKnownColumns()
	createPairKnownRows()

	//fmt.Println(len(rows))
	//fmt.Printf("parsing %v\n", time.Since(start))
	//start = time.Now()

	//pairs := findPairs(gal)
	//fmt.Printf("find pairs %v\n", time.Since(start))
	//fmt.Println(len(pairs))
	//start = time.Now()

	//fmt.Println(rows)
	//fmt.Println(columns)
	//printMap(b3)
	//printMapWithPoints(b3, gal)
	res1, res2 = sumDistance(gal)
	//fmt.Printf("find result %v\n", time.Since(start))
	//fmt.Println(result1)
	//fmt.Println(result2)
	return
}

func sumDistance(arr []point) (res1, res2 int64) {
	var x1, x2, y1, y2 int
	var r1, r2 int64
	lenA := len(arr)
	for i := 0; i < lenA; i++ {
		for k := i + 1; k < lenA; k++ {
			x1 = arr[i].x
			x2 = arr[k].x
			y1 = arr[i].y
			y2 = arr[k].y
			if arr[i].x > arr[k].x {
				x1, x2 = x2, x1
			}
			if arr[i].y > arr[k].y {
				y1, y2 = y2, y1
			}
			r1, r2 = sumDistanceBetweenPoints(x1, x2, y1, y2)
			res1 += r1
			res2 += r2
		}
	}
	return res1, res2
}

const part2Mul = 1_000_000 - 1

//const part2Mul = 99

func sumDistanceBetweenPoints(x1, x2, y1, y2 int) (int64, int64) {
	var part1, part2, tmp int
	part1 = columnMap[x1][x2]
	part2 = part1 * part2Mul

	tmp = rowMap[y1][y2]
	part1 += tmp
	part2 += tmp * part2Mul

	tmp = x2 - x1 + y2 - y1
	//fmt.Printf("len between %d and %d is %d\n", p[0].name, p[1].name, x+y+part1)
	return int64(tmp + part1), int64(tmp + part2)
}

func getExpanseBetween(from, to int, list []int) (res int) {
	if to-from == 1 {
		return 0
	}
	for i := from + 1; i < to; i++ {
		res += list[i]
	}
	return
}

func createPairs(arr []int) (res [][]int) {
	rowsCount := 0
	res = make([][]int, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		res = append(res, make([]int, 0, len(arr)))
		rowsCount = 0
		res[i] = append(res[i], rowsCount)
		for k := 0; k < len(arr); k++ {
			if k >= i+1 && arr[k] == 1 {
				rowsCount++
			}
			res[i] = append(res[i], rowsCount)
		}
	}
	return
}

func createPairKnownColumns() {
	rowsCount := 0
	for i := 0; i < 140; i++ {
		rowsCount = 0
		for k := 0; k < 140; k++ {
			if k >= i+1 && columns[k] == 1 {
				rowsCount++
			}
			if rowsCount == 0 {
				continue
			}
			columnMap[i][k] = rowsCount
		}
	}
	return
}

func createPairKnownRows() {
	rowCount := 0
	for i := 0; i < 140; i++ {
		rowCount = 0
		for k := 0; k < 140; k++ {
			if k >= i+1 && rows[k] == 1 {
				rowCount++
			}
			if rowCount == 0 {
				continue
			}
			rowMap[i][k] = rowCount
		}
	}
	return
}

func findPairs[A ~[]K, K any](arr A) (res []A) {
	res = make([]A, 0, 96141)
	for i := 0; i < len(arr); i++ {
		for k := i + 1; k < len(arr); k++ {
			e := make([]K, 2)
			e[0] = arr[i]
			e[1] = arr[k]
			res = append(res, e)
		}
	}
	return res
}

func createColumn(b3 [][]byte, i, offset int) {
	col := i + offset

	for y := 0; y < len(b3); y++ {
		m := make([]byte, len(b3[y])+1)
		copy(m, b3[y][:col])
		m[col] = '.'
		copy(m[col+1:], b3[y][col:])
		b3[y] = m
	}
}
func printMap(b3 [][]byte) {
	for y := 0; y < len(b3); y++ {
		fmt.Println(string(b3[y]))
	}
}

func printMapWithPoints(b3 [][]byte, gal []point) {

	for y := 0; y < len(b3); y++ {
		by := make([]byte, len(b3[y]))
		copy(by, b3[y])
		for i := 0; i < len(by); i++ {
			if by[i] == '#' {
				for _, p := range gal {
					if p.x == i && p.y == y {
						by[i] = ([]byte(strconv.Itoa(p.name)))[0]
						break
					}
				}
			}
		}
		fmt.Println(string(by))
	}
}

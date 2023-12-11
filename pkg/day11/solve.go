package day11

import (
	"bytes"
	"embed"
	"fmt"
	"strconv"
	"time"
)

//go:embed *.txt
var f embed.FS

type point struct {
	x, y, name int
}

func Solve() (err error) {
	var (
		b   []byte
		gal []point
	)
	b, err = f.ReadFile("input.txt")
	if err != nil {
		return
	}
	start := time.Now()

	tiles := bytes.Split(b, []byte{'\n'})
	rows := make([]int, len(tiles[0]))
	columns := make([]int, len(tiles[0]))
	galaxies := 0
	fmt.Printf("setup %v\n", time.Since(start))
	start = time.Now()
	for y := 0; y < len(tiles); y++ {
		galaxies = 0
		for x := 0; x < len(tiles[y]); x++ {
			if tiles[y][x] == '#' {
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
	fmt.Printf("parsing %v\n", time.Since(start))
	start = time.Now()

	//pairs := findPairs(gal)
	//fmt.Printf("find pairs %v\n", time.Since(start))
	//fmt.Println(len(pairs))
	//start = time.Now()

	//fmt.Println(rows)
	//fmt.Println(columns)
	//printMap(b3)
	//printMapWithPoints(b3, gal)
	result1, result2 := sumDistance(gal, rows, columns)
	fmt.Printf("find result %v\n", time.Since(start))
	fmt.Println(result1)
	fmt.Println(result2)
	return
}

func sumDistance(arr []point, rows, column []int) (res1 int, res2 int) {
	var r1, r2, x1, x2, y1, y2 int
	for i := 0; i < len(arr); i++ {
		for k := i + 1; k < len(arr); k++ {
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
			r1, r2 = sumDistanceBetweenPoints(x1, x2, y1, y2, rows, column)
			res1 += r1
			res2 += r2
		}
	}
	return res1, res2
}

const part2Mul = 1_000_000 - 1

//const part2Mul = 99

func sumDistanceBetweenPoints(x1, x2, y1, y2 int, rows, column []int) (int, int) {
	var part1, part2, tmp int
	part1 = getExpanseBetween(x1, x2, column)
	part2 = part1 * part2Mul

	tmp = getExpanseBetween(y1, y2, rows)
	part1 += tmp
	part2 += tmp * part2Mul

	tmp = x2 - x1 + y2 - y1
	//fmt.Printf("len between %d and %d is %d\n", p[0].name, p[1].name, x+y+part1)
	return tmp + part1, tmp + part2
}

func getExpanseBetween(from, to int, list []int) (res int) {
	for i := from + 1; i < to; i++ {
		res += list[i]
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

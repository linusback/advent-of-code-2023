package day11

import (
	"bytes"
	"embed"
	"fmt"
	"strconv"
)

//go:embed *.txt
var f embed.FS

type point struct {
	x, y, name int
}

func Solve() (err error) {
	var (
		b   []byte
		b3  [][]byte
		gal []point
	)
	b, err = f.ReadFile("input.txt")
	if err != nil {
		return
	}

	b2 := bytes.Split(b, []byte{'\n'})
	columns := make([]int, len(b2[0]))
	galaxies := 0
	//printMap(b2)
	for y := 0; y < len(b2); y++ {
		galaxies = 0
		b3 = append(b3, b2[y])
		for x := 0; x < len(b2[y]); x++ {
			if b2[y][x] == '#' {
				galaxies++
				columns[x]++
			}
		}
		if galaxies == 0 {
			b3 = append(b3, b2[y])
		}
	}
	i := 0
	//fmt.Println(columns)
	for x := 0; x < len(columns); x++ {
		if columns[x] > 0 {
			continue
		}
		//fmt.Println("create colum")
		createColumn(b3, x, i)
		i++
	}
	for y, row := range b3 {
		for x, c := range row {
			if c == '#' {
				gal = append(gal, point{
					x:    x,
					y:    y,
					name: len(gal) + 1,
				})
			}
		}
	}

	pairs := findPairs(gal)
	//printMap(b3)
	//printMapWithPoints(b3, gal)
	result1 := sumDistance(pairs)
	fmt.Println(result1)
	return
}

func sumDistance(pairs [][]point) int {
	var res int
	for _, pair := range pairs {
		res += sumDistanceBetweenPoints(pair)
	}
	return res
}

func sumDistanceBetweenPoints(p []point) int {
	var x, y int
	if p[0].x < p[1].x {
		x = p[1].x - p[0].x
	} else {
		x = p[0].x - p[1].x
	}

	if p[0].y < p[1].y {
		y = p[1].y - p[0].y
	} else {
		y = p[0].y - p[1].y
	}
	//fmt.Printf("len between %d and %d is %d\n", p[0].name, p[1].name, x+y)
	return x + y
}

func findPairs[A ~[]K, K any](arr A) (res []A) {
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

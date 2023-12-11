package day11

import (
	"embed"
	"io"
	"io/fs"
)

//go:embed *.txt
var f embed.FS

type point struct {
	x, y         uint64
	xExp, yExp   uint64
	xExp2, yExp2 uint64

	name int
}

const (
	galaxyByte byte   = '#'
	part2Mul   uint64 = 1_000_000 - 1
	//part2Mul uint64 = 99
)

func Solve() (res1, res2 int64, err error) {
	var (
		fil      fs.File
		gal      []point
		i, n     int
		galaxies uint8

		rowE, res1t, res2t, x uint64
	)
	fil, err = f.Open("input.txt")
	if err != nil {
		return
	}
	tiles := [140][140]byte{}
	columns := [140]uint8{}
	//tiles := [10][10]byte{}
	//columns := [10]uint8{}

	buff := make([]byte, 141)
	//buff := make([]byte, 11)

	gal = make([]point, 0, 440)

	for ; i < len(tiles); i++ {
		n, err = fil.Read(buff)
		if err != nil && err != io.EOF {
			return
		}
		if err == io.EOF {
			break
		}
		copy(tiles[i][:], buff[:n])
	}

	colGal := make([][]*point, len(tiles[0]))

	for z := 0; z < len(colGal); z++ {
		colGal[z] = make([]*point, 0, 6)
	}

	for y := uint64(0); y < uint64(len(tiles)); y++ {
		galaxies = 0
		for x = 0; x < uint64(len(tiles[y])); x++ {
			if tiles[y][x] == galaxyByte {
				columns[x]++
				galaxies++
				gal = append(gal, point{
					x:     x,
					y:     y,
					yExp:  y + rowE,
					yExp2: y + rowE*part2Mul,
					name:  len(gal) + 1,
				})
				res1t, res2t = calculateY(gal, res1t, res2t)
				colGal[x] = append(colGal[x], &gal[len(gal)-1])
			}
		}
		if galaxies == 0 {
			rowE++
		}
	}

	var colE uint64
	for z := 0; z < len(columns); z++ {
		if columns[z] > 0 {
			for i = 0; i < len(colGal[z]); i++ {

				colGal[z][i].xExp = colGal[z][i].x + colE
				colGal[z][i].xExp2 = colGal[z][i].x + colE*part2Mul
				res1t, res2t = calculateX(colGal[z][i], res1t, res2t, colGal[z][:i], colGal[:z]...)
			}
			continue
		}
		colE++
	}

	//res1t, res2t := sumDistance(gal)
	res1, res2 = int64(res1t), int64(res2t)
	return
}

func calculateY(gal []point, t, t2 uint64) (uint64, uint64) {
	var p, chosen *point
	chosen = &gal[len(gal)-1]
	for i := 0; i < len(gal)-1; i++ {
		p = &gal[i]
		t += chosen.yExp - p.yExp
		t2 += chosen.yExp2 - p.yExp2
	}
	return t, t2
}

func calculateX(chosen *point, t, t2 uint64, curr []*point, arr ...[]*point) (uint64, uint64) {
	var (
		p       *point
		j, lenJ int
	)

	for i := 0; i < len(arr); i++ {
		lenJ = len(arr[i])
		for j = 0; j < lenJ; j++ {
			p = arr[i][j]
			t += chosen.xExp - p.xExp
			t2 += chosen.xExp2 - p.xExp2
		}
	}
	for j = 0; j < len(curr); j++ {
		p = curr[j]
		t += chosen.xExp - p.xExp
		t2 += chosen.xExp2 - p.xExp2
	}
	return t, t2
}

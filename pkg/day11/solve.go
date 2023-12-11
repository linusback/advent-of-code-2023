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

		rowE, res1t, res2t, x, prev1, prev2 uint64
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
				res1t, res2t, prev1, prev2 = calculateY(gal, res1t, res2t, prev1, prev2)
				colGal[x] = append(colGal[x], &gal[len(gal)-1])
			}
		}
		if galaxies == 0 {
			rowE++
		}
	}

	prev1, prev2 = 0, 0
	x = 0
	rowE = 0
	for z := 0; z < len(columns); z++ {
		if columns[z] > 0 {
			for i = 0; i < len(colGal[z]); i++ {

				colGal[z][i].xExp = colGal[z][i].x + rowE
				colGal[z][i].xExp2 = colGal[z][i].x + rowE*part2Mul
				res1t, res2t, prev1, prev2 = calculateX(colGal[z][i], res1t, res2t, prev1, prev2, x)
				x++
			}
			continue
		}
		rowE++
	}

	res1, res2 = int64(res1t), int64(res2t)
	return
}

func calculateY(gal []point, t1, t2, prev1, prev2 uint64) (uint64, uint64, uint64, uint64) {
	var p *point
	mul := uint64(len(gal) - 1)
	p = &gal[mul]

	t1 += p.yExp*mul - prev1
	t2 += p.yExp2*mul - prev2

	prev1 += p.yExp
	prev2 += p.yExp2

	return t1, t2, prev1, prev2
}

func calculateX(p *point, t1, t2, prev1, prev2, mul uint64) (uint64, uint64, uint64, uint64) {
	t1 += p.xExp*mul - prev1
	t2 += p.xExp2*mul - prev2

	prev1 += p.xExp
	prev2 += p.xExp2

	return t1, t2, prev1, prev2
}

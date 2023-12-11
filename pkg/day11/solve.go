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
}

const (
	galaxyByte byte   = '#'
	part2Mul   uint64 = 1_000_000 - 1
	//part2Mul uint64 = 99
)

func Solve() (res1, res2 int64, err error) {
	var (
		fil      fs.File
		i        int
		galaxies uint8

		rowE, res1t, res2t, x, prev1, prev2, mul, e1, e2 uint64
	)
	fil, err = f.Open("input.txt")
	if err != nil {
		return
	}
	tiles := [140][141]byte{}
	columns := [140]uint8{}
	//tiles := [10][10]byte{}
	//columns := [10]uint8{}

	//buff := make([]byte, 141)
	//buff := make([]byte, 11)

	for ; i < len(tiles); i++ {
		_, err = fil.Read(tiles[i][:])
		if err != nil && err != io.EOF {
			return
		}
		if err == io.EOF {
			break
		}
	}

	colGal := make([]uint8, len(tiles[0]))

	for y := uint64(0); y < uint64(len(tiles)); y++ {
		galaxies = 0
		for x = 0; x < uint64(len(tiles[y])-1); x++ {
			if tiles[y][x] == galaxyByte {
				columns[x]++
				galaxies++
				e1 = y + rowE
				e2 = y + rowE*part2Mul
				colGal[x]++

				res1t, res2t, prev1, prev2 = calculate(e1, e2, res1t, res2t, prev1, prev2, mul)
				mul++
			}
		}
		if galaxies == 0 {
			rowE++
		}
	}

	prev1, prev2 = 0, 0
	x = 0
	mul = 0
	rowE = 0
	for ; x < uint64(len(columns)); x++ {
		if columns[x] > 0 {
			for galaxies = 0; galaxies < colGal[x]; galaxies++ {
				e1 = x + rowE
				e2 = x + rowE*part2Mul
				res1t, res2t, prev1, prev2 = calculate(e1, e2, res1t, res2t, prev1, prev2, mul)
				mul++
			}
			continue
		}
		rowE++
	}

	res1, res2 = int64(res1t), int64(res2t)
	return
}

func calculate(e1, e2, t1, t2, prev1, prev2, mul uint64) (uint64, uint64, uint64, uint64) {
	t1 += e1*mul - prev1
	t2 += e2*mul - prev2

	prev1 += e1
	prev2 += e2

	return t1, t2, prev1, prev2
}

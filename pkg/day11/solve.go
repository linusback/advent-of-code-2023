package day11

import (
	"embed"
	"io"
	"io/fs"
)

//go:embed *.txt
var f embed.FS

const (
	galaxyByte     byte   = '#'
	part2Mul       uint64 = 1_000_000 - 1
	rowLen, colLen        = 140, 140
)

func Solve() (res1, res2 int64, err error) {
	var (
		fil      fs.File
		galaxies uint8
		buff     [colLen + 1]byte
		columns  [colLen]uint8

		expanse1, expanse2, res1t, res2t, x, prev1, prev2, mul, e1, e2 uint64
	)

	fil, err = f.Open("input.txt")
	if err != nil {
		return
	}

	for y := uint64(0); y < rowLen; y++ {
		_, err = fil.Read(buff[:])
		if err != nil && err != io.EOF {
			return
		}
		galaxies = 0
		for x = 0; x < colLen; x++ {
			if buff[x] == galaxyByte {
				columns[x]++
				galaxies++
				e1 = y + expanse1
				e2 = y + expanse2

				res1t, res2t, prev1, prev2 = calculate(e1, e2, res1t, res2t, prev1, prev2, mul)
				mul++
			}
		}
		if galaxies == 0 {
			expanse1++
			expanse2 += part2Mul
		}
		if err == io.EOF {
			break
		}
	}

	prev1, prev2 = 0, 0
	x = 0
	mul = 0
	expanse1 = 0
	expanse2 = 0
	for ; x < colLen; x++ {
		if columns[x] > 0 {
			for galaxies = 0; galaxies < columns[x]; galaxies++ {
				e1 = x + expanse1
				e2 = x + expanse2
				res1t, res2t, prev1, prev2 = calculate(e1, e2, res1t, res2t, prev1, prev2, mul)
				mul++
			}
			continue
		}
		expanse1++
		expanse2 += part2Mul
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

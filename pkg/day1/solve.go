package day1

import (
	"advent-of-code-2023/pkg/util"
	"embed"
	"fmt"
)

//go:embed input.txt
//go:embed example2.txt
var f embed.FS

func Solve1() (err error) {
	var (
		b     []byte
		total int64
	)
	b, err = f.ReadFile("input.txt")
	if err != nil {
		return
	}

	err = util.DoEachRow(b, func(row []byte, nr int) error {
		var (
			start, end = -1, -1
		)
		for i := 0; i < len(row); i++ {
			if '0' <= row[i] && row[i] <= '9' {
				if start == -1 {
					start = i
				}
				end = i
			}
		}
		if start == -1 {
			return nil
		}
		total += int64((row[start] - '0') * 10)
		total += int64(row[end] - '0')
		return nil
	})
	fmt.Printf("total is %d\n", total)
	return
}

func Solve2() (err error) {
	var (
		b     []byte
		total int64
	)
	b, err = f.ReadFile("input.txt")
	if err != nil {
		return
	}

	//numbersAsText := map[string]int64{}

	err = util.DoEachRow(b, func(row []byte, nr int) error {
		var (
			start, end int64 = -1, -1
		)
		for i := 0; i < len(row); i++ {
			if '0' <= row[i] && row[i] <= '9' {
				if start == -1 {
					start = int64((row[i] - '0') * 10)
				}
				end = int64(row[i] - '0')
				continue
			}
			if res, _ := checkText(row[i:]); res > 0 {
				if start == -1 {
					start = res * 10
				}
				end = res
				//i += increment
			}

		}
		fmt.Printf("start: %d, end: %d, total: %d\n", start, end, start+end)
		fmt.Println("line: ", string(row))
		if start == -1 {
			return nil
		}
		total += start
		total += end
		return nil
	})
	fmt.Printf("total is %d\n", total)
	return
}

func checkText(sub []byte) (int64, int) {
	if len(sub) < 3 {
		return 0, 0
	}
	if len(sub) >= 5 {
		//fmt.Println("substring", string(sub[:5]))
		switch string(sub[:5]) {
		case "three":
			return 3, 4
		case "seven":
			return 7, 4
		case "eight":
			return 8, 4
		}
	}
	if len(sub) >= 4 {
		//fmt.Println("substring", string(sub[:4]))
		switch string(sub[:4]) {
		case "four":
			return 4, 3
		case "five":
			return 5, 3
		case "nine":
			return 9, 3
		}
	}
	//fmt.Println("substring", string(sub[:3]))
	switch string(sub[:3]) {
	case "one":
		return 1, 2
	case "two":
		return 2, 2
	case "six":
		return 6, 2
	}
	return 0, 0

}

package day6

import (
	"advent-of-code-2023/pkg/util"
	"embed"
	"fmt"
)

//go:embed *.txt
var f embed.FS

type Race struct {
	Time, Distance uint64
}

type Alternative struct {
	TimePressed, Distance uint64
}

func (r *Race) Alternatives() []Alternative {
	var d uint64
	res := make([]Alternative, 0, r.Time-2)
	for i := uint64(1); i < r.Time; i++ {
		d = i * (r.Time - i)
		if d > r.Distance {
			res = append(res, Alternative{
				TimePressed: i,
				Distance:    d,
			})
		}
	}
	return res
}

func (r *Race) AlternativesLen() (res uint64) {

	for i := uint64(1); i < r.Time; i++ {
		if i*(r.Time-i) > r.Distance {
			res++
		}
	}
	return res
}

func Solve() (err error) {
	var (
		b                []byte
		times, distances []uint64
		races            []Race
		res              = 1
		problem2         Race
	)
	b, err = f.ReadFile("example.txt")
	//b, err = f.ReadFile("input.txt")

	//p := util.NewTokenParser(b)
	//var r util.TokenSlice
	//for p.More() {
	//	r = p.NextRow()
	//	fmt.Println("name: ", string(r[0]))
	//	fmt.Println("val: ", r[1:].ToInt64())
	//}
	//return

	if err != nil {
		return
	}
	err = util.DoEachRow(b, func(row []byte, nr int) error {
		if nr == 0 {
			times = util.ParseUint64ArrNoError(row)
			problem2.Time = util.ParseUint64IgnoreAll(row)
			return nil
		}
		if nr == 1 {
			distances = util.ParseUint64ArrNoError(row)
			problem2.Distance = util.ParseUint64IgnoreAll(row)
			return nil
		}
		return nil
	})
	races = make([]Race, len(times))
	for i := 0; i < len(times); i++ {
		races[i].Time = times[i]
		races[i].Distance = distances[i]
		res *= len(races[i].Alternatives())
		//fmt.Println(races[i].Alternatives())
	}

	//fmt.Println(races)
	fmt.Println(problem2.AlternativesLen())
	fmt.Println(res)

	return
}

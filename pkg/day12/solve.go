package day12

import (
	"advent-of-code-2023/pkg/util"
	"embed"
	"fmt"
)

//go:embed *.txt
var f embed.FS

type springRow struct {
	springs      []byte
	springGroups []uint64
}

type groups [][]uint64

func Solve() (res1t, res2t uint64, err error) {
	var (
		b       []byte
		row     util.TokenSlice
		springs []springRow
	)
	b, err = f.ReadFile("input.txt")
	if err != nil {
		return
	}

	t := util.NewTokenParserSeparators(b, ' ', ',')
	springs = make([]springRow, 0, 1000)
	for t.More() {
		row = t.NextRow()
		s := springRow{
			springs: row[0],
		}
		combinations(row[0])
		s.springs = row[0]
		for i := 1; i < len(row); i++ {
			s.springGroups = append(s.springGroups, row[i].ParseUInt64())
		}
		fmt.Println(s.springGroups)
		springs = append(springs, s)
	}

	//fmt.Println(springs)
	return
}

func combinations(arr []byte) []groups {
	res := make([]groups, len(arr))
	insideGroup := false
	s := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == '.' {
			if insideGroup {
				insideGroup = false
				res = append(res, generateGroups(arr[s:i]))
			}

			continue
		}
		if !insideGroup {
			s = i
			insideGroup = true
		}
	}
	if insideGroup {
		res = append(res, generateGroups(arr[s:]))
	}
	return res

}

func generateGroups(bytes []byte) groups {
	res := make(groups, 0, len(bytes))
	fmt.Println("groups", string(bytes))
	return res
}

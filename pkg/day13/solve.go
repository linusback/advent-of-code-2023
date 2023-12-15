package day13

import (
	"advent-of-code-2023/pkg/util"
	"embed"
)

//go:embed *.txt
var f embed.FS

type Pattern struct {
	rows    [][]byte
	columns [][]byte
	//start of reflection since map 1 indexed
	rRow, rCol [2]int
	size       int
}

func (p *Pattern) Init(size int) {
	p.size = size
	p.rows = make([][]byte, 0, size)
	//p.columns = make([][]byte, size)
	//for i := 0; i < size; i++ {
	//	p.columns[i] = make([]byte, size)
	//}
}
func (p *Pattern) AddRow(row []byte) {
	//r := len(p.rows)
	p.rows = append(p.rows, row)
	//for i := 0; i < len(row); i++ {
	//	p.columns[i][r] = row[i]
	//}
}

func (p *Pattern) CheckColumns() {
	c := len(p.rows)
	r := len(p.rows[0])
	p.columns = make([][]byte, r)
	for i := 0; i < len(p.columns); i++ {
		p.columns[i] = make([]byte, c)
		for j := 0; j < len(p.columns[i]); j++ {
			p.columns[i][j] = p.rows[j][i]
		}
	}

	r1, r2 := CheckMirror(p.columns)
	if r1 != r2 {
		p.rCol[0] = r1
		p.rCol[1] = r2
	}
}
func (p *Pattern) CheckRows() {
	r1, r2 := CheckMirror(p.rows)
	if r1 != r2 {
		p.rRow[0] = r1
		p.rRow[1] = r2
	}
}

func CheckMirror(arr [][]byte) (r1 int, r2 int) {
	for i := 1; i < len(arr); i++ {
		if string(arr[i]) == string(arr[i-1]) {
			if isMirror(arr, i-1, i) {
				r1 = i - 1
				r2 = i
				return
			}
		}
	}
	return
}

func isMirror(arr [][]byte, up, down int) bool {
	for up >= 0 && down < len(arr) {
		if string(arr[up]) != string(arr[down]) {
			return false
		}
		up--
		down++
	}
	return true
}

func Solve() (res1t, res2t uint64, err error) {
	var (
		b []byte
	)
	b, err = f.ReadFile("input.txt")
	if err != nil {
		return
	}

	patterns := make([]Pattern, 0, 20)
	curr := Pattern{}
	err = util.DoEachRowAll(b, func(row []byte, rows [][]byte, nr, total int) error {
		if len(curr.rows) == 0 {
			curr.Init(len(row))
		}
		if len(row) == 0 {
			curr.CheckColumns()
			curr.CheckRows()
			patterns = append(patterns, curr)
			curr = Pattern{}
			return nil
		}
		curr.AddRow(row)

		return nil
	})
	curr.CheckColumns()
	curr.CheckRows()
	patterns = append(patterns, curr)

	for _, p := range patterns {
		if p.rCol[0] != p.rCol[1] {
			res1t += uint64(p.rCol[1])
		}
		if p.rRow[0] != p.rRow[1] {
			res1t += uint64(p.rRow[1] * 100)
		}
		//fmt.Println("row", p.rRow)
		//fmt.Println("col", p.rCol)
	}
	//fmt.Println(patterns)

	return
}

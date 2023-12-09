package day9

import (
	"advent-of-code-2023/pkg/util"
	"embed"
	"fmt"
)

//go:embed *.txt
var f embed.FS

func Solve() (err error) {
	var (
		b    []byte
		row  util.TokenSlice
		seq  [][]int64
		tree [][]int64
	)
	b, err = f.ReadFile("input.txt")
	if err != nil {
		return
	}
	t := util.NewTokenParser(b)
	longest := 0
	seq = make([][]int64, 0, 200)
	for t.More() {
		row = t.NextRow()
		s := make([]int64, len(row)+1)
		row.ConvertToInt64(s)
		if len(s) > longest {
			longest = len(s)
		}
		seq = append(seq, s)
	}
	tree = make([][]int64, longest-1)
	for i := 0; i < longest-1; i++ {
		tree[i] = make([]int64, longest-1-i)
	}
	var next, curr, s []int64
	var total, add, result1, result2 int64
	var i, j, k int
	resultArr := make([]int64, len(tree))
	for i = 0; i < len(seq); i++ {
		next = seq[i]
		s = seq[i]
		//fmt.Println(next)
		for j = 0; j < len(tree); j++ {
			total = 0
			curr = tree[j]
			for k = 0; k < len(s)-2-j; k++ {
				curr[k] = next[k+1] - next[k]
				total += curr[k]
			}
			if total == 0 {
				resultArr[j] = 0
				j--
				add = 0
				break
			}
			next = curr
			//fmt.Println(next)
		}

		for ; j >= 0; j-- {
			tree[j][len(s)-2-j] = tree[j][len(s)-3-j] + add
			add = tree[j][len(s)-2-j]
			resultArr[j] = tree[j][0] - resultArr[j+1]
			//fmt.Println(tree[j])
		}

		s[len(s)-1] = s[len(s)-2] + tree[0][len(s)-2]
		result1 += s[len(s)-1]
		result2 += s[0] - resultArr[0]
		//fmt.Println(s[0] - resultArr[0])
	}
	//fmt.Println(seq)
	fmt.Println(result1)
	fmt.Println(result2)
	return
}

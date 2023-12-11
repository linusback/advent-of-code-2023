package main

import (
	"advent-of-code-2023/pkg/day11"
	"fmt"
	"time"
)

func timeFunction(function func() (int64, int64, error)) {
	start := time.Now()
	res1, res2, err := function()
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("result 1: ", res1)
	fmt.Println("result 2: ", res2)

	fmt.Println("Time elapsed:", time.Since(start))
}
func main() {
	timeFunction(day11.Solve)
}

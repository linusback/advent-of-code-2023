package main

import (
	"advent-of-code-2023/pkg/day10"
	"fmt"
	"time"
)

func timeFunction(function func() error) {
	start := time.Now()
	err := function()
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("Time elapsed:", time.Since(start))
}
func main() {
	timeFunction(day10.Solve)
}

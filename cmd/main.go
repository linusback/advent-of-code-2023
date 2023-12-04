package main

import (
	"advent-of-code-2023/pkg/day2"
	"fmt"
	"log"
)

func main() {
	fmt.Println("hello advent of code 2023")
	err := day2.Solve1()
	if err != nil {
		log.Fatal(err)
	}
}

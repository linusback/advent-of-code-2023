package main

import (
	"advent-of-code-2023/pkg/day1"
	"fmt"
	"log"
)

func main() {
	fmt.Println("hello advent of code 2023")
	err := day1.Solve1()
	if err != nil {
		log.Fatal(err)
	}
	err = day1.Solve2()
	if err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"fmt"
	"os"
	"time"

	"aoc/internal/utils"
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day06.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	start := time.Now()

	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s\n", elapsed)
	return 0
}

// part two
func part2(input string) int {
	start := time.Now()

	elapsed := time.Since(start)
	fmt.Printf("Part 2 took %s\n", elapsed)
	return 0
}

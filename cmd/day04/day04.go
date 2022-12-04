package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"aoc/internal/utils"
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/dayXX.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// // read from http request
	// session := os.Getenv("AOC_SESSION")
	// input, err := utils.ReadHTTP(2021, 1, session)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(0)
	// }

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// ███████╗██╗   ██╗██████╗ ██████╗  ██████╗ ██████╗ ████████╗
// ██╔════╝██║   ██║██╔══██╗██╔══██╗██╔═══██╗██╔══██╗╚══██╔══╝
// ███████╗██║   ██║██████╔╝██████╔╝██║   ██║██████╔╝   ██║
// ╚════██║██║   ██║██╔═══╝ ██╔═══╝ ██║   ██║██╔══██╗   ██║
// ███████║╚██████╔╝██║     ██║     ╚██████╔╝██║  ██║   ██║
// ╚══════╝ ╚═════╝ ╚═╝     ╚═╝      ╚═════╝ ╚═╝  ╚═╝   ╚═╝

func splitOnComma(input string) []string {
	return strings.Split(input, ",")
}

func splitOnDash(input string) []string {
	return strings.Split(input, "-")
}

// ██████╗  █████╗ ██████╗ ████████╗ ██╗
// ██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝███║
// ██████╔╝███████║██████╔╝   ██║   ╚██║
// ██╔═══╝ ██╔══██║██╔══██╗   ██║    ██║
// ██║     ██║  ██║██║  ██║   ██║    ██║
// ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝    ╚═╝

// part one
func part1(input string) int {
	start := time.Now()

	for {
		line, err := buffer.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s\n", elapsed)
	return 0
}

// part two
func part2(input string) int {
	start := time.Now()

	for {
		line, err := buffer.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Part 2 took %s\n", elapsed)
	return 0
}

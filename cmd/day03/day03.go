package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"aoc/internal/utils"
)

func main() {
	input, err := utils.ReadBuffer("resources/day03.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))

	input, err = utils.ReadBuffer("resources/day03.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
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

func createRunesAsBinary() map[rune]int {
	// create a map from rune to int representation where the rune is a-Z
	// and the int representation is 2^0, 2^1, 2^2, etc.
	var runesAsBinary = make(map[rune]int)

	for i, r := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		runesAsBinary[r] = 1 << i // 2^i
	}

	return runesAsBinary
}

func createPriorities() map[int]int {
	// create a map from int to int where the int is 2^0, 2^1, 2^2, etc.
	// and the int is the priority 1, 2, 3, etc.
	var priorities = make(map[int]int)

	for i, _ := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		priorities[1<<i] = i + 1 // 2^i
	}

	return priorities
}

func reduce(input string, runeMap map[rune]int) int {
	// reduce the input string to a single int value
	// by ORing all the rune values together
	var result int

	for _, r := range input {
		result |= runeMap[r]
	}

	return result
}

// ██████╗  █████╗ ██████╗ ████████╗ ██╗
// ██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝███║
// ██████╔╝███████║██████╔╝   ██║   ╚██║
// ██╔═══╝ ██╔══██║██╔══██╗   ██║    ██║
// ██║     ██║  ██║██║  ██║   ██║    ██║
// ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝    ╚═╝

func part1(buffer *bufio.Reader) int {
	start := time.Now()
	var result int

	runesAsBinary := createRunesAsBinary()
	priorities := createPriorities()

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

		halfSize := len(line) / 2
		leftSack := line[:halfSize]
		rightSack := line[halfSize : len(line)-1]

		leftSackReduced := reduce(leftSack, runesAsBinary)
		rightSackReduced := reduce(rightSack, runesAsBinary)
		filteredSacks := leftSackReduced & rightSackReduced
		priority := priorities[filteredSacks]

		result += priority
	}

	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s\n", elapsed)
	return result
}

// ██████╗  █████╗ ██████╗ ████████╗██████╗
// ██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝╚════██╗
// ██████╔╝███████║██████╔╝   ██║    █████╔╝
// ██╔═══╝ ██╔══██║██╔══██╗   ██║   ██╔═══╝
// ██║     ██║  ██║██║  ██║   ██║   ███████╗
// ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚══════╝

func part2(buffer *bufio.Reader) int {
	start := time.Now()
	var result, counter, accumulated int
	// allPriorities is 2^52-1, which is all the bits set to 1
	var allPriorities = (1 << 52) - 1
	accumulated = allPriorities

	runesAsBinary := createRunesAsBinary()
	priorities := createPriorities()

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

		counter++
		accumulated &= reduce(line, runesAsBinary)

		if counter%3 == 0 {
			result += priorities[accumulated]
			accumulated = allPriorities
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s\n", elapsed)
	return result
}

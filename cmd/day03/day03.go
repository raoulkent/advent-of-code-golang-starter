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

func runePriority(r rune) int {
	if int(r) >= 97 {
		return 1 << (int(r) - 97)
	} else {
		return 1 << (int(r) - 65 + 26)
	}
}

func priorityRune(i int) rune {
	if i < 26 {
		return rune(i + 97)
	} else {
		return rune(i + 65 - 26)
	}
}

func msb(n int) int {
	var pos int
	for n > 0 {
		n >>= 1
		pos++
	}
	return pos
}

func reduce(input string) int {
	var result int
	for _, r := range input {
		result |= runePriority(r)
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

		leftSackReduced := reduce(leftSack)
		rightSackReduced := reduce(rightSack)
		filteredSacks := leftSackReduced & rightSackReduced
		priority := msb(filteredSacks)

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
	var allPriorities = (1 << 52) - 1
	accumulated = allPriorities

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
		accumulated &= reduce(line[:len(line)-1])

		if counter%3 == 0 {
			result += msb(accumulated)
			accumulated = allPriorities
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s\n", elapsed)
	return result
}

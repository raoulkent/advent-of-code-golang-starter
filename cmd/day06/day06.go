package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"aoc/internal/utils"
)

func main() {
	var input1, input2 *bufio.Reader

	fmt.Println("--- Part One ---")
	input1, _ = utils.ReadBuffer("resources/day06.txt")
	fmt.Println("Result:", part1(input1))

	fmt.Println("--- Part Two ---")
	input2, _ = utils.ReadBuffer("resources/day06.txt")
	fmt.Println("Result:", part2(input2))

	os.Exit(0)
}

// ███████╗██╗   ██╗██████╗ ██████╗  ██████╗ ██████╗ ████████╗
// ██╔════╝██║   ██║██╔══██╗██╔══██╗██╔═══██╗██╔══██╗╚══██╔══╝
// ███████╗██║   ██║██████╔╝██████╔╝██║   ██║██████╔╝   ██║
// ╚════██║██║   ██║██╔═══╝ ██╔═══╝ ██║   ██║██╔══██╗   ██║
// ███████║╚██████╔╝██║     ██║     ╚██████╔╝██║  ██║   ██║
// ╚══════╝ ╚═════╝ ╚═╝     ╚═╝      ╚═════╝ ╚═╝  ╚═╝   ╚═╝

func bytesAreUnique(bytes []byte) bool {
	seen := make(map[byte]bool)
	for _, r := range bytes {
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

// ██████╗  █████╗ ██████╗ ████████╗     ██████╗ ███╗   ██╗███████╗
// ██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝    ██╔═══██╗████╗  ██║██╔════╝
// ██████╔╝███████║██████╔╝   ██║       ██║   ██║██╔██╗ ██║█████╗
// ██╔═══╝ ██╔══██║██╔══██╗   ██║       ██║   ██║██║╚██╗██║██╔══╝
// ██║     ██║  ██║██║  ██║   ██║       ╚██████╔╝██║ ╚████║███████╗
// ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝        ╚═════╝ ╚═╝  ╚═══╝╚══════╝

func part1(buffer *bufio.Reader) int {
	start := time.Now()

	// The message begins after 4 characters (the header of the message)
	var position int = 0 + 4

	for {
		peek, _ := buffer.Peek(4)
		if bytesAreUnique(peek) {
			fmt.Printf("Found unique bytes at position %d: %s\n", position, peek)
			break
		}

		buffer.ReadRune()
		position++
	}

	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s\n", elapsed)
	return position
}

// ██████╗  █████╗ ██████╗ ████████╗    ████████╗██╗    ██╗ ██████╗
// ██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝    ╚══██╔══╝██║    ██║██╔═══██╗
// ██████╔╝███████║██████╔╝   ██║          ██║   ██║ █╗ ██║██║   ██║
// ██╔═══╝ ██╔══██║██╔══██╗   ██║          ██║   ██║███╗██║██║   ██║
// ██║     ██║  ██║██║  ██║   ██║          ██║   ╚███╔███╔╝╚██████╔╝
// ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝          ╚═╝    ╚══╝╚══╝  ╚═════╝

func part2(buffer *bufio.Reader) int {
	start := time.Now()

	elapsed := time.Since(start)
	fmt.Printf("Part 2 took %s\n", elapsed)
	return 0
}

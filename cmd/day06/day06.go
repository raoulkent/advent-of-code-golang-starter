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

func findBufferedBytesUntilUnique(buffer *bufio.Reader, syncLength int) int {
	// The message begins after syncLength number of characters (the sync header of the message)
	var position int = 0 + syncLength

	//TODO: Use a scanner, with ScanRunes to read the buffer instead
	//TODO 2: Instead of peeking, read the buffer and store the last four runes in a queue

	for {
		peek, _ := buffer.Peek(syncLength)
		if bytesAreUnique(peek) {
			break
		}

		buffer.ReadRune()
		position++
	}

	return position
}

// ██████╗  █████╗ ██████╗ ████████╗     ██████╗ ███╗   ██╗███████╗
// ██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝    ██╔═══██╗████╗  ██║██╔════╝
// ██████╔╝███████║██████╔╝   ██║       ██║   ██║██╔██╗ ██║█████╗
// ██╔═══╝ ██╔══██║██╔══██╗   ██║       ██║   ██║██║╚██╗██║██╔══╝
// ██║     ██║  ██║██║  ██║   ██║       ╚██████╔╝██║ ╚████║███████╗
// ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝        ╚═════╝ ╚═╝  ╚═══╝╚══════╝

func part1(buffer *bufio.Reader) int {
	start := time.Now()

	processedRunes := findBufferedBytesUntilUnique(buffer, 4)

	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s\n", elapsed)
	return processedRunes
}

// ██████╗  █████╗ ██████╗ ████████╗    ████████╗██╗    ██╗ ██████╗
// ██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝    ╚══██╔══╝██║    ██║██╔═══██╗
// ██████╔╝███████║██████╔╝   ██║          ██║   ██║ █╗ ██║██║   ██║
// ██╔═══╝ ██╔══██║██╔══██╗   ██║          ██║   ██║███╗██║██║   ██║
// ██║     ██║  ██║██║  ██║   ██║          ██║   ╚███╔███╔╝╚██████╔╝
// ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝          ╚═╝    ╚══╝╚══╝  ╚═════╝

func part2(buffer *bufio.Reader) int {
	start := time.Now()

	processedRunes := findBufferedBytesUntilUnique(buffer, 14)

	elapsed := time.Since(start)
	fmt.Printf("Part 2 took %s\n", elapsed)
	return processedRunes
}

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
	input1, _ = utils.ReadBuffer("resources/day05.txt")
	fmt.Println("Result:", part1(input1))

	fmt.Println("--- Part Two ---")
	input2, _ = utils.ReadBuffer("resources/day05.txt")
	fmt.Println("Result:", part2(input2))

	os.Exit(0)
}

// ____ _  _ ___  ___  ____ ____ ___
// [__  |  | |__] |__] |  | |__/  |
// ___] |__| |    |    |__| |  \  |

// Place support functions here

// ___  ____ ____ ___    ____ _  _ ____
// |__] |__| |__/  |     |  | |\ | |___
// |    |  | |  \  |     |__| | \| |___

func part1(buffer *bufio.Reader) int {
	start := time.Now()

	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s\n", elapsed)
	return 0
}

// ___  ____ ____ ___    ___ _ _ _ ____
// |__] |__| |__/  |      |  | | | |  |
// |    |  | |  \  |      |  |_|_| |__|

func part2(buffer *bufio.Reader) int {
	start := time.Now()

	elapsed := time.Since(start)
	fmt.Printf("Part 2 took %s\n", elapsed)
	return 0
}

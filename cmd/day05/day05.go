package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gammazero/deque"

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

func inputParser(buffer *bufio.Reader) []string {
	var initialStackState []string

	for {
		line, err := buffer.ReadString('\n')
		if err == io.EOF {
			break
		}
		if isMoveString(line) {
			break
		}

		initialStackState = append(initialStackState, line)
	}

	return initialStackState
}

func isMoveString(s string) bool {
	return strings.Split(s, " ")[0] == "move"
}

func startStateParser(buffer *bufio.Reader) []deque.Deque[rune] {
	var q []deque.Deque[rune]

	return q
}

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

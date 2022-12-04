package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"aoc/internal/utils"
)

func main() {
	// read form file
	buffer, err := utils.ReadBuffer("resources/day04.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(buffer))

	buffer, err = utils.ReadBuffer("resources/day04.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(buffer))

	os.Exit(0)
}

func part1(buffer *bufio.Reader) int {
	start := time.Now()

	var count int

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

		line = strings.TrimSuffix(line, "\n")

		pairs := strings.Split(line, ",")
		var lower, upper []int
		for _, pair := range pairs {
			split := strings.Split(pair, "-")

			leftInt, _ := strconv.Atoi(split[0])
			rightInt, _ := strconv.Atoi(split[1])

			lower = append(lower, leftInt)
			upper = append(upper, rightInt)
		}

		if lower[0] <= lower[1] && upper[0] >= upper[1] {
			count++
		} else if lower[0] >= lower[1] && upper[0] <= upper[1] {
			count++
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Part 1 took %s\n", elapsed)
	return count
}

func part2(buffer *bufio.Reader) int {
	start := time.Now()

	var count int

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

		line = strings.TrimSuffix(line, "\n")

		pairs := strings.Split(line, ",")
		var lower, upper []int
		for _, pair := range pairs {
			split := strings.Split(pair, "-")

			leftInt, _ := strconv.Atoi(split[0])
			rightInt, _ := strconv.Atoi(split[1])

			lower = append(lower, leftInt)
			upper = append(upper, rightInt)
		}

		if lower[0] <= lower[1] && lower[1] <= upper[0] {
			count++
		} else if lower[1] <= lower[0] && lower[0] <= upper[1] {
			count++
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Part 2 took %s\n", elapsed)
	return count
}

package main

import (
	"fmt"
	"os"
	"strings"

	"aoc/internal/utils"
)

func main() {
	input, err := utils.ReadFile("resources/day02.txt")
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

// ███████╗██╗   ██╗██████╗ ██████╗  ██████╗ ██████╗ ████████╗
// ██╔════╝██║   ██║██╔══██╗██╔══██╗██╔═══██╗██╔══██╗╚══██╔══╝
// ███████╗██║   ██║██████╔╝██████╔╝██║   ██║██████╔╝   ██║
// ╚════██║██║   ██║██╔═══╝ ██╔═══╝ ██║   ██║██╔══██╗   ██║
// ███████║╚██████╔╝██║     ██║     ╚██████╔╝██║  ██║   ██║
// ╚══════╝ ╚═════╝ ╚═╝     ╚═╝      ╚═════╝ ╚═╝  ╚═╝   ╚═╝

var outcomeMap = map[string]map[string]int{
	"X": {
		"A": 3,
		"B": 0,
		"C": 6,
	},
	"Y": {
		"A": 6,
		"B": 3,
		"C": 0,
	},
	"Z": {
		"A": 0,
		"B": 6,
		"C": 3,
	},
}

var shapeMap = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func splitOnNewlines(s string) []string {
	removeEOF := strings.TrimSuffix(s, "\n")
	return strings.Split(removeEOF, "\n")
}

func splitOnSpace(s string) []string {
	return strings.Split(s, " ")
}

func getShapeScore(shape string) int {
	return shapeMap[shape]
}

func getOutcomeScore(shape string, elfShape string) int {
	return outcomeMap[shape][elfShape]
}

// ██████╗  █████╗ ██████╗ ████████╗     ██╗
// ██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝    ███║
// ██████╔╝███████║██████╔╝   ██║       ╚██║
// ██╔═══╝ ██╔══██║██╔══██╗   ██║        ██║
// ██║     ██║  ██║██║  ██║   ██║        ██║
// ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝        ╚═╝

func part1(input string) int {
	var score int = 0

	rounds := splitOnNewlines(input)

	for _, round := range rounds {
		roundData := splitOnSpace(round)

		elfShape := roundData[0]
		shape := roundData[1]

		shapeScore := getShapeScore(shape)
		outcomeScore := getOutcomeScore(shape, elfShape)

		score += (shapeScore + outcomeScore)
	}

	return score
}

// ██████╗  █████╗ ██████╗ ████████╗    ██████╗
// ██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝    ╚════██╗
// ██████╔╝███████║██████╔╝   ██║        █████╔╝
// ██╔═══╝ ██╔══██║██╔══██╗   ██║       ██╔═══╝
// ██║     ██║  ██║██║  ██║   ██║       ███████╗
// ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝       ╚══════╝

func part2(input string) int {
	return 0
}

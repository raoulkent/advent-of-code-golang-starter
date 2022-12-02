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
	"A": {
		"X": 3,
		"Y": 0,
		"Z": 6,
	},
	"B": {
		"X": 6,
		"Y": 3,
		"Z": 0,
	},
	"C": {
		"X": 0,
		"Y": 6,
		"Z": 3,
	},
}

var shapeMap = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

func splitOnNewlines(s string) []string {
	return strings.Split(s, "\n")
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

		// fmt.Printf("Round: %s\n", roundData)
		// fmt.Printf("Shape: %s\n", roundData[0])
		// fmt.Printf("Elfshape: %s\n", roundData[1])
		shape := roundData[0]
		elfShape := roundData[1]

		// shapeScore := getShapeScore(shape)
		outcomeScore := getOutcomeScore(shape, elfShape)

		// score += (shapeScore + outcomeScore)
		score += outcomeScore
		// score += shapeScore
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

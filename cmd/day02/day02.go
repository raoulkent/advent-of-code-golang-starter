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

var outcomes = map[string]map[string]int{
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

var shapes = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var secretShape = map[string]map[string]string{
	"X": {
		"A": "Z",
		"B": "X",
		"C": "Y",
	},
	"Y": {
		"A": "X",
		"B": "Y",
		"C": "Z",
	},
	"Z": {
		"A": "Y",
		"B": "Z",
		"C": "X",
	},
}

var secretOutcome = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

func splitOnNewlines(s string) []string {
	removeEOF := strings.TrimSuffix(s, "\n")
	return strings.Split(removeEOF, "\n")
}

func splitOnSpace(s string) []string {
	return strings.Split(s, " ")
}

func getShapeScore(shape string) int {
	return shapes[shape]
}

func getOutcomeScore(elfShape string, shape string) int {
	return outcomes[elfShape][shape]
}

func getSecretShape(elfShape string, secretInstruction string) string {
	return secretShape[secretInstruction][elfShape]
}

func getSecretOutcome(shape string) int {
	return secretOutcome[shape]
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
	var score int = 0

	rounds := splitOnNewlines(input)

	for _, round := range rounds {
		roundData := splitOnSpace(round)

		elfShape := roundData[0]          // X, Y or Z
		secretInstruction := roundData[1] // A, B, or C
		shape := getSecretShape(elfShape, secretInstruction)

		fmt.Printf("Elf: %s, Shape: %s, Secret Shape: %s\n", elfShape, secretInstruction, shape)

		shapeScore := getShapeScore(shape)
		outcomeScore := getSecretOutcome(secretInstruction)

		score += (shapeScore + outcomeScore)
	}

	return score
}

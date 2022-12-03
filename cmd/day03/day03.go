package main

import (
	"fmt"
	"os"

	"aoc/internal/utils"
)

func main() {
	input, err := utils.ReadFile("resources/day03.txt")
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

func createRunesAsBinary() map[rune]int {
	// create a map from rune to int representation where the rune is a-Z
	// and the int representation is 2^0, 2^1, 2^2, etc.
	var runesAsBinary = make(map[rune]int)

	for i, r := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		runesAsBinary[r] = 1 << i // 2^i
	}

	return runesAsBinary
}

func createBinaryAsRunes() map[int]rune {
	// create a map from int representation to rune where the int representation is 2^0, 2^1, 2^2, etc.
	// and the rune is a-Z
	var binaryAsRunes = make(map[int]rune)

	for i, r := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		binaryAsRunes[1<<i] = r // 2^i
	}

	return binaryAsRunes
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

func part1(input string) int {
	// runesAsBinary := createRunesAsBinary()
	// binaryAsRunes := createBinaryAsRunes()
	return 0
}

// ██████╗  █████╗ ██████╗ ████████╗██████╗
// ██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝╚════██╗
// ██████╔╝███████║██████╔╝   ██║    █████╔╝
// ██╔═══╝ ██╔══██║██╔══██╗   ██║   ██╔═══╝
// ██║     ██║  ██║██║  ██║   ██║   ███████╗
// ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚══════╝

func part2(input string) int {
	return 0
}

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"aoc/internal/utils"
)

func main() {
	input, err := utils.ReadFile("resources/day01.txt")
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

type elfCandyBag struct {
	elf        int
	candyTotal int
}

type elfCandyBags []elfCandyBag

func (e elfCandyBags) Len() int {
	return len(e)
}

func (e elfCandyBags) Less(i, j int) bool {
	return e[i].candyTotal < e[j].candyTotal
}

func (e elfCandyBags) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// ██████╗  █████╗ ██████╗ ████████╗     ██████╗ ███╗   ██╗███████╗
// ██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝    ██╔═══██╗████╗  ██║██╔════╝
// ██████╔╝███████║██████╔╝   ██║       ██║   ██║██╔██╗ ██║█████╗
// ██╔═══╝ ██╔══██║██╔══██╗   ██║       ██║   ██║██║╚██╗██║██╔══╝
// ██║     ██║  ██║██║  ██║   ██║       ╚██████╔╝██║ ╚████║███████╗
// ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝        ╚═════╝ ╚═╝  ╚═══╝╚══════╝

func part1(input string) int {
	candyBags := bagStringsToElfCandyBags(input)

	sort.Sort(sort.Reverse(candyBags))
	fmt.Printf("Elf with the most candy: %d \n", candyBags[0].elf)

	return candyBags[0].candyTotal
}

func bagStringsToElfCandyBags(input string) elfCandyBags {
	var candyBags elfCandyBags

	bagStrings := strings.Split(input, "\n\n")

	for i, s := range bagStrings {
		candy := strings.Split(s, "\n")
		candySum := sumArr(convertArrStringToArrInt(candy))

		// Add the elf i, and the candyTotal candySum to the candyBags
		candyBags = append(candyBags, elfCandyBag{i, candySum})
	}

	return candyBags
}

func convertArrStringToArrInt(input []string) []int {
	var output []int

	for _, v := range input {
		value, _ := strconv.Atoi(v)
		output = append(output, value)
	}

	return output
}

func sumArr(input []int) int {
	var sum int

	for _, v := range input {
		sum += v
	}

	return sum
}

// ██████╗  █████╗ ██████╗ ████████╗    ████████╗██╗    ██╗ ██████╗
// ██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝    ╚══██╔══╝██║    ██║██╔═══██╗
// ██████╔╝███████║██████╔╝   ██║          ██║   ██║ █╗ ██║██║   ██║
// ██╔═══╝ ██╔══██║██╔══██╗   ██║          ██║   ██║███╗██║██║   ██║
// ██║     ██║  ██║██║  ██║   ██║          ██║   ╚███╔███╔╝╚██████╔╝
// ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝          ╚═╝    ╚══╝╚══╝  ╚═════╝

func part2(input string) int {
	candyBags := bagStringsToElfCandyBags(input)

	sort.Sort(sort.Reverse(candyBags))

	var candySumForTop3Elf int

	for i := 0; i < 3; i++ {
		candySumForTop3Elf += candyBags[i].candyTotal
	}

	return candySumForTop3Elf
}

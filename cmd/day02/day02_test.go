package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay02(t *testing.T) {
	assert := assert.New(t)
	input := "A Y\nB X\nC Z\n"

	t.Run("part 1", func(t *testing.T) {
		expected := 15
		actual := part1(input)

		fmt.Printf("Part 1: %d", actual)
		assert.Equal(actual, expected)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 0
		actual := part2(input)

		assert.Equal(actual, expected)
	})
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	input := "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw"

	t.Run("part 1", func(t *testing.T) {
		expected := 157
		actual := part1(input)

		assert.Equal(actual, expected)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 0
		actual := part2(input)

		assert.Equal(actual, expected)
	})
}

package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayXX(t *testing.T) {
	assert := assert.New(t)
	input := bufio.NewReader(strings.NewReader(""))

	t.Run("part 1", func(t *testing.T) {
		expected := 0
		actual := part1(input)

		assert.Equal(actual, expected)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 0
		actual := part2(input)

		assert.Equal(actual, expected)
	})
}

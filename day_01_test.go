package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	test_input := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
	expected_result := []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}
	result, err := parseInput(test_input)

	assert.Equal(t, result, expected_result)
	assert.Equal(t, err, nil)

}

func TestRotateTest(t *testing.T) {
	actions := []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}
	start_position := 50
	expected_result := 3
	result := rotate_count_zero_pos(start_position, actions)

	assert.Equal(t, result, expected_result)
}

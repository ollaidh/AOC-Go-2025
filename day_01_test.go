package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	testInput := []string{"L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}
	expectedResult := []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}
	result, err := parseInput(testInput)

	assert.Equal(t, result, expectedResult)
	assert.Equal(t, err, nil)

}

func TestRotateCountZeroPos(t *testing.T) {
	actions := []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}
	startPosition := 50
	expectedResult := 3
	result := rotateCountZeroPos(startPosition, actions)

	assert.Equal(t, expectedResult, result)
}

func TestRotateCountAllZeroes(t *testing.T) {
	actions := []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}
	startPosition := 50
	expectedResult := 6
	result := rotateCountAllZeroes(startPosition, actions)

	fmt.Printf("%d", result)

	assert.Equal(t, expectedResult, result)
}

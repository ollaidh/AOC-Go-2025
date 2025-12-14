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

func TestStep(t *testing.T) {
	assert.Equal(t, 82, step(50, -68).position)
	assert.Equal(t, 52, step(82, -30).position)
	assert.Equal(t, 0, step(52, 48).position)
	assert.Equal(t, 95, step(0, -5).position)
	assert.Equal(t, 55, step(95, 60).position)
	assert.Equal(t, 0, step(55, -55).position)
	assert.Equal(t, 99, step(0, -1).position)
	assert.Equal(t, 0, step(99, -99).position)
	assert.Equal(t, 14, step(0, 14).position)
	assert.Equal(t, 32, step(14, -82).position)
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

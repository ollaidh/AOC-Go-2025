package main

import (
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

func TestStepPosition(t *testing.T) {
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

func TestStepZeroesCount(t *testing.T) {
	assert.Equal(t, 1, step(50, -68).zeroesCount)
	assert.Equal(t, 0, step(82, -30).zeroesCount)
	assert.Equal(t, 1, step(52, 48).zeroesCount)
	assert.Equal(t, 0, step(0, -5).zeroesCount)
	assert.Equal(t, 1, step(95, 60).zeroesCount)
	assert.Equal(t, 1, step(55, -55).zeroesCount)
	assert.Equal(t, 0, step(0, -1).zeroesCount)
	assert.Equal(t, 1, step(99, -99).zeroesCount)
	assert.Equal(t, 0, step(0, 14).zeroesCount)
	assert.Equal(t, 1, step(14, -82).zeroesCount)

	assert.Equal(t, 10, step(50, 1000).zeroesCount)
	assert.Equal(t, 10, step(50, -1000).zeroesCount)
	assert.Equal(t, 10, step(50, 999).zeroesCount)
	assert.Equal(t, 10, step(50, -999).zeroesCount)
}

func TestRotateCountZeroPos(t *testing.T) {
	actions := []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}
	result := rotateCountZeroPos(50, actions)
	assert.Equal(t, 3, result)
}

func TestRotateCountAllZeroes(t *testing.T) {
	actions := []int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82}
	result := rotateCountAllZeroes(50, actions)
	assert.Equal(t, 6, result)
}

package main

import (
	"os"
	"strconv"
	"strings"
)

type Day1Part1 struct{}

func getInputDay1(inputRaw string) ([]int, error) {
	input := strings.Split(inputRaw, "\n")
	result := []int{}
	for _, action := range input {
		value, err := strconv.Atoi(action[1:])
		if err != nil {
			return nil, err
		}
		if action[:1] == "L" {
			value = value * (-1)
		}
		result = append(result, value)
	}
	return result, nil
}

func (day Day1Part1) solve(rawInput string) string {
	actions, err := getInputDay1(rawInput)
	if err != nil {
		panic(err)
	}
	resultPart1 := rotateCountZeroPos(50, actions)
	return strconv.Itoa(resultPart1)

}

type Day1Part2 struct{}

func (day Day1Part2) solve(rawInput string) string {
	actions, err := getInputDay1(rawInput)
	if err != nil {
		panic(err)
	}
	resultPart2 := rotateCountAllZeroes(50, actions)
	return strconv.Itoa(resultPart2)

}

type StepResult struct {
	position    int
	zeroesCount int
}

func getInputFromFile(inputPath string) string {
	b, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func step(startPosition int, step int) StepResult {
	fullCirclesCount := max(step, -step) / 100

	step = step % 100
	newPosition := (startPosition + 100 + step) % 100

	if step < 0 && -step >= startPosition && startPosition != 0 {
		fullCirclesCount++
	}
	if step > 0 && newPosition < startPosition {
		fullCirclesCount++
	}

	return StepResult{
		position:    newPosition,
		zeroesCount: fullCirclesCount,
	}
}

func rotateCountZeroPos(startPosition int, actions []int) int {
	result := 0

	for _, action := range actions {
		currResult := step(startPosition, action)
		if currResult.position == 0 {
			result++
		}
		startPosition = currResult.position
	}
	return result
}

func rotateCountAllZeroes(startPosition int, actions []int) int {
	result := 0

	for _, action := range actions {
		currResult := step(startPosition, action)
		result += currResult.zeroesCount
		startPosition = currResult.position
	}
	return result
}

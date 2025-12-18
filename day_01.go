package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Day1Part1 struct{}

func (day Day1Part1) solve([]string) string {
	inputPath := getInputFilePath("/inputs/input_day_01-1.dat")
	input := getLinesFromFile(inputPath)
	actions, err := parseInput(input)
	if err != nil {
		panic(err)
	}

	resultPart1 := rotateCountZeroPos(50, actions)
	fmt.Println(resultPart1)

	resultPart2 := rotateCountAllZeroes(50, actions)
	fmt.Println(resultPart2)

}

type Day1Part2 struct{}

func (day Day1Part2) solve([]string) string {
}

type StepResult struct {
	position    int
	zeroesCount int
}

func getLinesFromFile(inputPath string) []string {
	f, _ := os.Open(inputPath)
	scanner := bufio.NewScanner(f)
	result := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

func parseInput(input []string) ([]int, error) {
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

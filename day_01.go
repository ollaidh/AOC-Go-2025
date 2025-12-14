package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func rotateCountZeroPos(startPosition int, actions []int) int {
	result := 0

	for _, step := range actions {
		step = step % 100
		startPosition = startPosition + step

		if startPosition < 0 {
			startPosition = startPosition + 100
		} else if startPosition > 99 {
			startPosition = startPosition - 100
		}
		if startPosition == 0 {
			result++
		}
	}
	return result
}

func rotateCountAllZeroes(startPosition int, actions []int) int {
	result := 0

	for _, step := range actions {
		additionalZeroes := step / 100
		if additionalZeroes < 0 {
			additionalZeroes = additionalZeroes * (-1)
		}
		fmt.Printf("%v %d \n", step, additionalZeroes)

		step = step % 100
		newPosition := startPosition + step
		fmt.Printf("step %d, pos %d, new_pos %v ", step, startPosition, newPosition)

		if newPosition < 0 {
			newPosition = newPosition + 100
		} else if newPosition > 99 {
			newPosition = newPosition - 100
		}

		if step <= 0 && startPosition <= newPosition {
			result++
		}
		if step > 0 && startPosition > newPosition {
			result++
		}

		result = result + additionalZeroes
		startPosition = newPosition
		fmt.Printf("%d \n", result)

	}
	return result
}

func main() {
	inputPath := getInputFilePath("/inputs/input_day_01-1.dat")
	input := getLinesFromFile(inputPath)
	actions, err := parseInput(input)
	if err != nil {
		panic(err)
	}

	resultPart1 := rotateCountZeroPos(50, actions)
	fmt.Printf("%d", resultPart1)

	resultPart2 := rotateCountAllZeroes(50, actions)
	fmt.Printf("%d", resultPart2)

}

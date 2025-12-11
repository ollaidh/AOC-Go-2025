package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getInputFromFile(inputPath string) []string {
	f, _ := os.Open(inputPath)
	scanner := bufio.NewScanner(f)
	result := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

func parseInput(input []string) []int {
	result := []int{}
	for _, action := range input {
		value, _ := strconv.Atoi(action[1:])
		if action[:1] == "L" {
			value = value * (-1)
		}
		result = append(result, value)
	}
	return result
}

func main() {
	inputPath := get_input_file_path("/inputs/input_day_01-1.dat")
	input := getInputFromFile(inputPath)
	actions := parseInput(input)
	curr_position := 50
	result := 0

	for _, step := range actions {
		step = step % 100
		curr_position = curr_position + step

		if curr_position < 0 {
			curr_position = curr_position + 100
		} else if curr_position > 99 {
			curr_position = curr_position - 100
		}
		if curr_position == 0 {
			result++
		}
	}
	fmt.Printf("%d", result)

}

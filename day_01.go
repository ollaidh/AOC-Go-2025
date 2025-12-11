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

func main() {
	inputPath := get_input_file_path("/inputs/input_day_01-1.dat")
	input := getLinesFromFile(inputPath)
	actions, err := parseInput(input)
	if err != nil {
		panic(err)
	}
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseInput(inputPath string) [][]string {
	f, _ := os.Open(inputPath)
	scanner := bufio.NewScanner(f)
	actions := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[:1]
		steps := line[1:]
		one_action := []string{direction, steps}
		actions = append(actions, one_action)
	}
	return actions
}

func main() {
	inputPath := get_input_file_path("/inputs/input_day_01-1.dat")
	actions := parseInput(inputPath)
	curr_position := 50
	result := 0

	for _, v := range actions {
		steps, _ := strconv.Atoi(v[1])
		steps = steps % 100

		if v[0] == "L" {
			curr_position = curr_position - steps
		} else {
			curr_position = curr_position + steps
		}

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

package main

import (
	"fmt"
)

type solver interface {
	solve([]string) string
}

type Date struct {
	year int
	day  int
	part int
}

// TODO add iteratinf over all versions of input for this day

func main() {
	tasks := map[Date]solver{
		Date{year: 2025, day: 1, part: 1}: Day1Part1{},
		Date{year: 2025, day: 1, part: 2}: Day1Part2{},
	}
	var result string

	inputsFolderPath := getInputsFolderPath()

	for date, callable := range tasks {
		filePaths := getInputFilePaths(date, inputsFolderPath)

		for _, filepath := range filePaths {
			input := getDataFromFile(filepath)
			result = callable.solve(input)
			fmt.Printf("Year %v Day %v Part %v Path %v result: %v\n", date.year, date.day, date.part, filepath, result)
		}
	}

}

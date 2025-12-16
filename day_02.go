package main

import (
	"strconv"
	"strings"
)

type Range struct {
	minValue int
	maxValue int
}

func parseInputDayTwo(input string) []Range {
	result := []Range{}
	inputRanges := strings.Split(input, ",")
	for _, currRange := range inputRanges {
		currCurrRange := strings.Split(currRange, "-")
		start, _ := strconv.Atoi(currCurrRange[0])
		end, _ := strconv.Atoi(currCurrRange[1])
		result = append(result, Range{minValue: start, maxValue: end})
	}
	return result
}

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInputDayTwo(t *testing.T) {
	initialInput := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	result := parseInputDayTwo(initialInput)
	assert.Equal(t, 11, len(result))
	assert.Equal(t, 11, result[0].minValue)
	assert.Equal(t, 22, result[0].maxValue)
	assert.Equal(t, 95, result[1].minValue)
	assert.Equal(t, 115, result[1].maxValue)
	assert.Equal(t, 998, result[2].minValue)
	assert.Equal(t, 1012, result[2].maxValue)
	assert.Equal(t, 1188511880, result[3].minValue)
	assert.Equal(t, 1188511890, result[3].maxValue)
	assert.Equal(t, 222220, result[4].minValue)
	assert.Equal(t, 222224, result[4].maxValue)
	assert.Equal(t, 1698522, result[5].minValue)
	assert.Equal(t, 1698528, result[5].maxValue)
	assert.Equal(t, 446443, result[6].minValue)
	assert.Equal(t, 446449, result[6].maxValue)
	assert.Equal(t, 38593856, result[7].minValue)
	assert.Equal(t, 38593862, result[7].maxValue)
	assert.Equal(t, 565653, result[8].minValue)
	assert.Equal(t, 565659, result[8].maxValue)
	assert.Equal(t, 824824821, result[9].minValue)
	assert.Equal(t, 824824827, result[9].maxValue)
	assert.Equal(t, 2121212118, result[10].minValue)
	assert.Equal(t, 2121212124, result[10].maxValue)

}

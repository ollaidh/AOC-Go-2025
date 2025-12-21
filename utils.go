package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func getInputFileName(date Date) string {
	filename := fmt.Sprintf("input_day_%s_part_%s.dat", strconv.Itoa(date.day), strconv.Itoa(date.part))
	return filename
}

func getInputFilePath(filename string) string {. // TODO do func with env variable
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	var s strings.Builder
	s.WriteString(basepath + filename)
	inputPath := s.String()
	return inputPath
}

func getDataFromFile(date Date) []string {
	filename := getInputFileName(date)
	inputFolderPath := getInputFilePath(filename). // TODO func not working

	inputPath := fmt.Sprintf("%s/%s", inputFolderPath, filename)

	f, _ := os.Open(inputPath)
	scanner := bufio.NewScanner(f)
	result := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

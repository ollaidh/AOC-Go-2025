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

func getInputFilePaths(date Date, filesPath string) []string {
	fileName := fmt.Sprintf("input_day_%s_part_%s.dat", strconv.Itoa(date.day), strconv.Itoa(date.part))
	fileNames := []string{fileName}
	for i := range fileNames {
		fileNames[i] = fmt.Sprintf("%s/%s", filesPath, fileNames[i])
	}
	return fileNames
}

func getInputsFolderPath() string { // TODO do func with env variable
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	var s strings.Builder
	s.WriteString(basepath)
	inputPath := s.String()
	return inputPath
}

func getDataFromFile(inputPath string) []string {
	f, _ := os.Open(inputPath)
	scanner := bufio.NewScanner(f)
	result := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

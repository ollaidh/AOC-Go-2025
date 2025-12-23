package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func getInputsFolderPath() (string, error) {
	inputsFolderRootPath := os.Getenv("INPUTS_FOLDER")
	if inputsFolderRootPath == "" {
		return "", errors.New("No INPUTS_FOLDER environment variable is found!")
	}
	return inputsFolderRootPath, nil
}

func getInputFilePaths(date Date, filesPath string) []string {
	fileName := fmt.Sprintf("input_day_%s_part_%s.dat", strconv.Itoa(date.day), strconv.Itoa(date.part))
	fileNames := []string{fileName}
	for i := range fileNames {
		fileNames[i] = filepath.Join(filesPath, fileNames[i])
	}
	return fileNames
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

package main

import (
	"path/filepath"
	"runtime"
	"strings"
)

func getInputFilePath(relativePath string) string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	var s strings.Builder
	s.WriteString(basepath + relativePath)
	inputPath := s.String()
	return inputPath
}

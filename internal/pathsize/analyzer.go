// Package analyzer provides functions for calculating the size of files
// and directories, with support for recursive traversal and hidden files.
package pathsize

import (
	"os"
	"path/filepath"
	"strings"
	"fmt"
)

type CLIArgs struct {
	Recursive bool
	All       bool
	Human     bool
}

// AnalyzeFile returns the size of a single file at the given path.
func analyzeFile(cliArgs CLIArgs, path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	return info.Size(), nil
}



func isHidden(path string) bool {
	base := filepath.Base(path)

	return strings.HasPrefix(base, ".")
}

// porcessFolder process one DirEntry.
func porcessFolder(cliArgs CLIArgs, path string, entry os.DirEntry) (int64, error) {
	fullPath := filepath.Join(path, entry.Name())

	if !cliArgs.All && isHidden(fullPath) {
		return 0, nil
	}

	if entry.IsDir() {
		if !cliArgs.Recursive {
			return 0, nil
		}

		return AnalyzeFolder(cliArgs, fullPath)
	}

	fileInfo, err := entry.Info()
	if err != nil {
		return 0, err
	}

	return fileInfo.Size(), nil
}

// AnalyzeFolder returns the size of a folder at the given path.
func analyzeFolder(cliArgs CLIArgs, path string) (int64, error) {
	var totalSize int64

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		size, err := porcessFolder(cliArgs, path, entry)
		if err != nil {
			return totalSize, err
		}

		totalSize += size
	}

	return totalSize, nil
}

// Analyze path size.
func Analyze(cliArgs CLIArgs, path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}


	var result int64
	if !info.IsDir() {
		result, err = analyzeFile(cliArgs, path)
	} else {
		result, err = analyzeFolder(cliArgs, path)
	}

	if err != nil {
		return 0, err
	}

	return result, nil
}



const (
	bytesInKB   = 1024
	roundFactor = 100
)

var sizeSuffixes = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

func humanize(size int64) string {
	sizeSuffixesIndex := 0

	if size < int64(bytesInKB) {
		return fmt.Sprintf("%dB", size)
	}

	result := float64(size)
	for result >= float64(bytesInKB) {
		result /= bytesInKB
		sizeSuffixesIndex++
	}


	if result == float64(int64(result)) {
		return fmt.Sprintf("%d.0%s", int64(result), sizeSuffixes[sizeSuffixesIndex])
	}

	return fmt.Sprintf("%.2f%s", result, sizeSuffixes[sizeSuffixesIndex])
}


func GetHumanFormattedResult(cliArgs CLIArgs, size int64, path string) string {
	var formattedResult string
	
	if cliArgs.Human {
		formattedResult = humanize(size)
	} else {
		formattedResult = fmt.Sprintf("%dB", size)
	}

	if path != "" {
		return fmt.Sprintf("%s\t%s\n", formattedResult, path)
	}

	return formattedResult
	
}

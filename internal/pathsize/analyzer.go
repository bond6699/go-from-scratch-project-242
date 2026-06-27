// Package analyzer provides functions for calculating the size of files
// and directories, with support for recursive traversal and hidden files.
package pathsize

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CLIArgs struct {
	Recursive bool
	All       bool
	Human     bool
}

// sizeSuffixes list all aviable size suffixes
var sizeSuffixes = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

const bytesInKB int64 = 1024

// humanizeSize size(int64) -> format output(string)
func humanizeSize(size int64) string {
	sizeSuffixesIndex := 0

	if size < bytesInKB {
		return fmt.Sprintf("%dB", size)
	}

	result := float64(size)
	for result >= float64(bytesInKB) {
		result /= float64(bytesInKB)
		sizeSuffixesIndex++
	}

	return fmt.Sprintf("%.2f%s", result, sizeSuffixes[sizeSuffixesIndex])
}

// GetHumanFormattedResult wrap humanizeSize with cliArgs & path
func GetHumanFormattedResult(cliArgs CLIArgs, size int64, path string) string {
	var formattedResult string

	if cliArgs.Human {
		formattedResult = humanizeSize(size)
	} else {
		formattedResult = fmt.Sprintf("%dB", size)
	}

	if path != "" {
		return fmt.Sprintf("%s\t%s\n", formattedResult, path)
	}

	return formattedResult
}

// analyzeFile returns the size of a single file at the given path.
func analyzeFile(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	return info.Size(), nil
}

func IsSymLink(path string) (bool, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return true, err
	}

	return fileInfo.Mode()&os.ModeSymlink != 0, nil
}

func isHidden(path string) bool {
	base := filepath.Base(path)

	return strings.HasPrefix(base, ".")
}

// porcessFolder process one DirEntry.
func porcessFolder(cliArgs CLIArgs, path string, entry os.DirEntry) (int64, error) {
	fullPath := filepath.Join(path, entry.Name())

	isSumlink, err := IsSymLink(fullPath)
	if err != nil {
		return 0, err
	}

	if isSumlink {
		return 0, nil
	}

	if !cliArgs.All && isHidden(fullPath) {
		return 0, nil
	}

	if entry.IsDir() {
		if !cliArgs.Recursive {
			return 0, nil
		}

		return analyzeFolder(cliArgs, fullPath)
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
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		isSumlink, err := IsSymLink(path)
		if err != nil {
			return 0, err
		}

		if isSumlink {
			return 0, nil
		}

		return analyzeFile(path)
	} else {
		return analyzeFolder(cliArgs, path)
	}
}

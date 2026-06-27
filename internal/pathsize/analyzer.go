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

func IsSymLink(path string) bool {
	fileInfo, _ := os.Lstat(path)

	return fileInfo.Mode()&os.ModeSymlink != 0
}

func isHidden(path string) bool {
	base := filepath.Base(path)

	return strings.HasPrefix(base, ".")
}

//nolint:gocognit,gocyclo,gofumpt // Complexity is acceptable due to multiple skip conditions
func analyzeFolder(cliArgs CLIArgs, rootPath string) (int64, error) {
	var totalSize int64

	err := filepath.WalkDir(rootPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("cannot access %s: %w", path, err)
		}

		if IsSymLink(path) {
			return nil
		}

		if !cliArgs.All && isHidden(path) {
			if d.IsDir() {
				return filepath.SkipDir
			}

			return nil
		}

		if d.IsDir() && !cliArgs.Recursive && path != rootPath {
			return filepath.SkipDir
		}

		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return fmt.Errorf("cannot get file info for %s: %w", path, err)
			}

			totalSize += info.Size()
		}

		return nil
	})

	return totalSize, err
}

// Analyze path size.
func Analyze(cliArgs CLIArgs, path string) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		if IsSymLink(path) {
			return 0, nil
		}

		return analyzeFile(path)
	} else {
		return analyzeFolder(cliArgs, path)
	}
}

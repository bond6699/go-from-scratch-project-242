// Package analyzer provides functions for calculating the size of files
// and directories, with support for recursive traversal and hidden files.
package pathsize

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Options struct {
	Recursive bool
	All       bool
	Human     bool
	Path      string
}

// sizeSuffixes list all aviable size suffixes
var sizeSuffixes = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

const bytesInKB int64 = 1024

// humanizeSize size(int64) -> format output(string)
func HumanizeSize(size int64) string {
	sizeSuffixesIndex := 0

	if size < bytesInKB {
		return fmt.Sprintf("%dB", size)
	}

	// потому что так требуют тесты

	result := float64(size)
	for result >= float64(bytesInKB) {
		result /= float64(bytesInKB)
		sizeSuffixesIndex++
	}

	if result == float64(int64(result)) {
		return fmt.Sprintf("%d.0%s", int64(result), sizeSuffixes[sizeSuffixesIndex])
	}

	return fmt.Sprintf("%.2f%s", result, sizeSuffixes[sizeSuffixesIndex])
}

// GetHumanFormattedResult wrap humanizeSize with cliArgs & path
func GetHumanFormattedResult(human bool, size int64, path string) string {
	var formattedResult string

	if human {
		formattedResult = HumanizeSize(size)
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

func isHidden(path string) bool {
	base := filepath.Base(path)

	return strings.HasPrefix(base, ".")
}

//nolint:gocognit,gocyclo,gofumpt // Complexity is acceptable due to multiple skip conditions
func analyzeFolder(recursive, all bool, rootPath string) (int64, error) {
	var totalSize int64

	err := filepath.WalkDir(rootPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("cannot access %s: %w", path, err)
		}

		realPath, err := filepath.EvalSymlinks(path)
		if err != nil {
			return err
		}

		fmt.Println(realPath, path)

		if realPath != path {
			path = realPath
		}

		if !all && isHidden(path) {
			if d.IsDir() {
				return filepath.SkipDir
			}

			return nil
		}

		if d.IsDir() && !recursive && path != rootPath {
			return filepath.SkipDir
		}

		if !d.IsDir() {
			info, err := os.Stat(path)
			if err != nil {
				return fmt.Errorf("cannot get file info for %s: %w", path, err)
			}

			totalSize += info.Size()
			fmt.Println(realPath, path, info.Size())
		}

		return nil
	})

	return totalSize, err
}

// Analyze path size.
func Analyze(recursive, all bool, path string) (int64, error) {
	realPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		return 0, err
	}

	if realPath != path {
		path = realPath
	}

	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return analyzeFile(path)
	} else {
		return analyzeFolder(recursive, all, path)
	}
}

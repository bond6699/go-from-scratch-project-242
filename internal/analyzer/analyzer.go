// Package analyzer provides functions for calculating the size of files
// and directories, with support for recursive traversal and hidden files.
package analyzer

import (
	"code/internal/flags"
	"os"
	"path/filepath"
	"strings"
)

func isHidden(path string) bool {
	base := filepath.Base(path)

	return strings.HasPrefix(base, ".")
}

// AnalyzeFile returns the size of a single file at the given path.
func AnalyzeFile(cmdFlags flags.CLIFlags, path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	if !cmdFlags.All && isHidden(path) {
		return 0, nil
	}

	return info.Size(), nil
}

// porcessFolder process one DirEntry.
func porcessFolder(cmdFlags flags.CLIFlags, path string, entry os.DirEntry) (int64, error) {
	fullPath := filepath.Join(path, entry.Name())

	if !cmdFlags.All && isHidden(fullPath) {
		return 0, nil
	}

	if entry.IsDir() {
		if !cmdFlags.Recursive {
			return 0, nil
		}

		return AnalyzeFolder(cmdFlags, fullPath)
	}

	fileInfo, err := entry.Info()
	if err != nil {
		return 0, err
	}

	return fileInfo.Size(), nil
}

// AnalyzeFolder returns the size of a folder at the given path.
func AnalyzeFolder(cmdFlags flags.CLIFlags, path string) (int64, error) {
	var totalSize int64

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, entry := range entries {
		size, err := porcessFolder(cmdFlags, path, entry)
		if err != nil {
			return totalSize, err
		}

		totalSize += size
	}

	return totalSize, nil
}

// Analyze path size.
func Analyze(cmdFlags flags.CLIFlags, path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return AnalyzeFile(cmdFlags, path)
	}

	return AnalyzeFolder(cmdFlags, path)
}

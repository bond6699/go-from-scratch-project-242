// Package analyzer provides functions for calculating the size of files
// and directories, with support for recursive traversal and hidden files.
package analyzer

import (
	"os"
	"path/filepath"

	"code/internal/flags"
)

func isHidden(path string) bool {
	base := filepath.Base(path)

	return len(base) > 0 && string(base[0]) == "."
}

// AnalyzeFile returns the size of a single file at the given path.
func AnalyzeFile(flags flags.CLIFlags, path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	if !flags.All && isHidden(path) {
		return 0, nil
	}

	return info.Size(), nil
}

// AnalyzeFolder returns the size of a folder at the given path.
func AnalyzeFolder(flags flags.CLIFlags, path string) (int64, error) {
	var totalSize int64

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, e := range entries {
		fullPath := filepath.Join(path, e.Name())

		// Пропускаем скрытые, если All выключен
		if !flags.All && isHidden(fullPath) {
			continue
		}

		if e.IsDir() {
			if !flags.Recursive {
				continue
			}

			size, err := AnalyzeFolder(flags, fullPath)
			if err != nil {
				return totalSize, err
			}

			totalSize += size

			continue
		}

		fileInfo, err := e.Info()
		if err != nil {
			return totalSize, err
		}

		totalSize += fileInfo.Size()
	}

	return totalSize, nil
}

// Analyze path size.
func Analyze(flags flags.CLIFlags, path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return AnalyzeFile(flags, path)
	}

	return AnalyzeFolder(flags, path)
}

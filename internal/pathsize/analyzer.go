// Package analyzer provides functions for calculating the size of files
// and directories, with support for recursive traversal and hidden files.
package pathsize

import (
	internal_errors "code/internal/errors"
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

// analyzeFile returns the size of a single file at the given path.
func analyzeFile(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, fmt.Errorf("%w: %s %w", internal_errors.ErrReadPath, path, err)
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
			return fmt.Errorf("%w: %s %w", internal_errors.ErrAccessDenied, path, err)
		}

		realPath, err := filepath.EvalSymlinks(path)
		if err != nil {
			return internal_errors.ErrSymslinkRead
		}

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
				return fmt.Errorf("%w: %s %w", internal_errors.ErrReadPath, path, err)
			}

			totalSize += info.Size()
		}

		return nil
	})

	return totalSize, err
}

// Analyze path size.
func Analyze(recursive, all bool, path string) (int64, error) {
	if realPath, err := filepath.EvalSymlinks(path); err == nil {
		path = realPath
	}

	info, err := os.Stat(path)
	if err != nil {
		return 0, fmt.Errorf("%w: %s %w", internal_errors.ErrReadPath, path, err)
	}

	if !info.IsDir() {
		return analyzeFile(path)
	} else {
		return analyzeFolder(recursive, all, path)
	}
}

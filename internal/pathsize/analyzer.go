// Package analyzer provides functions for calculating the size of files
// and directories, with support for recursive traversal and hidden files.
package pathsize

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Options struct {
	Recursive bool
	All       bool
	Path      string
}

func isHidden(path string) bool {
	return strings.HasPrefix(filepath.Base(path), ".")
}

//nolint:gocognit,gocyclo,gofumpt // Complexity is acceptable due to multiple skip conditions
func analyzeFolder(options Options) (int64, error) {
	var totalSize int64

	rootPath, err := filepath.EvalSymlinks(options.Path)
	if err != nil {
		return 0, mapPathError(rootPath, err)
	}

	err = filepath.WalkDir(rootPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return mapPathError(path, err)
		}

		if !options.All && isHidden(path) {
			if d.IsDir() {
				return filepath.SkipDir
			}

			return nil
		}

		if d.IsDir() && !options.Recursive && path != rootPath {
			return filepath.SkipDir
		}

		if !d.IsDir() {
			info, err := d.Info()

			if d.Type()&fs.ModeSymlink != 0 {
				info, err = os.Stat(path)
			}

			if err != nil {
				return mapPathError(path, err)
			}

			totalSize += info.Size()
		}

		return nil
	})

	return totalSize, err
}

func Analyze(options Options) (int64, error) {
	info, err := os.Stat(options.Path)
	if err != nil {
		return 0, mapPathError(options.Path, err)
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	return analyzeFolder(options)
}

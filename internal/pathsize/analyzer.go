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
func analyzeFolder(recursive, all bool, rootPath string) (int64, error) {
	var totalSize int64

	rootPath, err := filepath.EvalSymlinks(rootPath)
	if err != nil {
		return 0, mapPathError(rootPath, err)
	}

	err = filepath.WalkDir(rootPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return mapPathError(path, err)
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

func Analyze(recursive, all bool, path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, mapPathError(path, err)
	}

	if !info.IsDir() {
		return info.Size(), nil
	} 
	
	return analyzeFolder(recursive, all, path)
}

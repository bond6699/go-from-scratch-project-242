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
}

func isHidden(path string) bool {
	return strings.HasPrefix(filepath.Base(path), ".")
}

//nolint:gocognit,gocyclo,gofumpt // Complexity is acceptable due to multiple skip conditions
func getFolderSize(options Options, path string) (int64, error) {
	var totalSize int64

	rootPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		return 0, getPathError(path, err)
	}

	err = filepath.WalkDir(rootPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return getPathError(path, err)
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

		mode := d.Type()

		if mode&fs.ModeSymlink != 0 {
			info, err := os.Stat(path)
			if err != nil {
				return getPathError(path, err)
			}

			totalSize += info.Size()

			return nil
		}

		if mode&fs.ModeType == 0 {
			info, err := d.Info()
			if err != nil {
				return getPathError(path, err)
			}

			totalSize += info.Size()
		}

		return nil
	})

	return totalSize, err
}

func GetPathSize(options Options, path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, getPathError(path, err)
	}

	if info.IsDir() {
		return getFolderSize(options, path)
	}

	return info.Size(), nil
}

package pathsize

import (
	"errors"
	"fmt"
	"io/fs"
)

var (
	errInvalidPath  = errors.New("invalid path")
	errAccessDenied = errors.New("access denied")
	errReadPath     = errors.New("cannot read path")
)

func mapPathError(path string, err error) error {
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return fmt.Errorf("%w: %s: %w", errInvalidPath, path, err)
	case errors.Is(err, fs.ErrPermission):
		return fmt.Errorf("%w: %s: %w", errAccessDenied, path, err)
	default:
		return fmt.Errorf("%w: %s: %w", errReadPath, path, err)
	}
}

package code

import (
	"code/internal/formatter"
	"code/internal/pathsize"
)

// GetPathSize возвращает человекочитаемый размер пути в виде строки.
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := pathsize.GetPathSize(pathsize.Options{
		Recursive: recursive,
		All:       all,
	}, path)
	if err != nil {
		return "", err
	}

	formattedSize := formatter.FormatSize(human, size)

	return formattedSize, nil
}

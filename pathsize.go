package code

import (
	"code/internal/formatter"
	"code/internal/pathsize"
)

// GetPathSize возвращает человекочитаемый размер пути в виде строки.
func GetPathSize(path string, recursive, all bool) (string, error) {
	result, err := pathsize.Analyze(pathsize.Options{
		Recursive: recursive,
		All:       all,
		Path:      path,
	})
	if err != nil {
		return "", err
	}

	formattedResult := formatter.HumanizeSize(result, true)

	return formattedResult, nil
}

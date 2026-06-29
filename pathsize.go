package code

import (
	"code/internal/pathsize"
	"code/internal/formatter"
)

// GetPathSize возвращает человекочитаемый размер пути в виде строки.
func GetPathSize(path string, recursive, all bool) (string, error) {
	result, err := pathsize.Analyze(recursive, all, path)
	if err != nil {
		return "", err
	}

	formattedResult := formatter.HumanizeSize(result)

	return formattedResult, nil
}

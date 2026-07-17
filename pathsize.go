package code

import (
	"code/internal/formatter"
	"code/internal/pathsize"
)

// GetPathSize возвращает человекочитаемый размер пути в виде строки.
func GetPathSize(path string, recursive, all, human bool) (string, error) {
	result, err := pathsize.GetPathSize(pathsize.Options{
		Recursive: recursive,
		All:       all,
	}, path)
	if err != nil {
		return "", err
	}

	formattedResult := formatter.GetHumanResult(human, result, "")

	return formattedResult, nil
}

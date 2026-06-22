package code

import (
	"code/internal/analyzer"
	"code/internal/flags"
	"code/internal/formatter"
	"fmt"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	cliflags := flags.Create(recursive, human, all)

	size, err := analyzer.Analyze(cliflags, path)
	if err != nil {
		return "", err
	}

	sizeStr := formatter.Formatter(cliflags, size)

	return fmt.Sprintf("%s\t%s", sizeStr, path), nil
}

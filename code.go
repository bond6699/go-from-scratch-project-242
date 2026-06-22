package code

import (
	"code/internal/analyzer"
	"code/internal/flags"
)

func GetPathSize(path string, recursive, human, all bool) (int64, error) {
	cliflags := flags.Create(recursive, human, all)

	result, err := analyzer.Analyze(cliflags, path)
	if err != nil {
		return 0, err
	}

	return result, nil
}

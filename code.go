package code

import (
	"code/internal/analyzer"
	"code/internal/flags"
)

func GetPathSize(path string) (int64, error) {
	cliflags := flags.Create(true, true, true)

	result, err := analyzer.Analyze(cliflags, path)
	if err != nil {
		return 0, err
	}

	return result, nil
}

package code

import (
	"code/internal/analyzer"
	"code/internal/flags"
)

func GetPathSize(path string) (int64, error) {
	cliflags := flags.Create(true, true, true)
	return analyzer.Analyzer()
}
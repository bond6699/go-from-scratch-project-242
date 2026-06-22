package code

import (
	"code/internal/analyzer"
	"code/internal/flags"
	"code/internal/formatter"
)

// GetPathSize возвращает человекочитаемый размер пути в виде строки.
// Рекурсивно обрабатывает директории, игнорирует скрытые файлы.
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	// Устанавливаем флаги: recursive=true, human=true, all=false
	cliflags := flags.Create(recursive, all, human)

	size, err := analyzer.Analyze(cliflags, path)
	if err != nil {
		return "", err
	}

	sizeStr := formatter.Formatter(cliflags, size)

	return sizeStr, nil
}

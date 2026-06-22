package code

import (
	"code/internal/analyzer"
	"code/internal/flags"
	"code/internal/formatter"
	"code/internal/cli"
	"fmt"
)

// GetPathSize возвращает человекочитаемый размер пути в виде строки.
// Рекурсивно обрабатывает директории, игнорирует скрытые файлы.
func GetPathSize(path string) (string, error) {
	// Устанавливаем флаги: recursive=true, human=true, all=false
	cliflags := flags.Create(true, true, false)

	size, err := analyzer.Analyze(cliflags, path)
	if err != nil {
		return "", err
	}

	sizeStr := formatter.Formatter(cliflags, size)
	return sizeStr, nil
}

func main() {
	cli.Run()
}

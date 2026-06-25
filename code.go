package code

import (
	"code/internal/pathsize"
)

// GetPathSize возвращает человекочитаемый размер пути в виде строки.
// Рекурсивно обрабатывает директории, игнорирует скрытые файлы.
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	// Устанавливаем флаги: recursive=true, human=true, all=false
	cliArgs := pathsize.CLIArgs{recursive, all, human}

	result, err := pathsize.Analyze(cliArgs, path)
	if err != nil {
		return "", err
	}

	formattedResult := pathsize.GetFormattedResult(cliArgs, result, path)

	return formattedResult, nil
}

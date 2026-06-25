package code

import (
	"code/internal/pathsize"
	"fmt"
)

// GetPathSize возвращает человекочитаемый размер пути в виде строки.
// Рекурсивно обрабатывает директории, игнорирует скрытые файлы.
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	// Устанавливаем флаги: recursive=true, human=true, all=false
	cliArgs := pathsize.Create(recursive, all, human)

	result, err := pathsize.Analyze(cliArgs, path)
	if err != nil {
		return "", err
	}

	var formattedResult string
	if cliArgs.Human {
		formattedResult = pathsize.Humanity(result)
	} else {
		formattedResult = fmt.Sprintf("%dB", result)
	}
	

	return formattedResult, nil
}

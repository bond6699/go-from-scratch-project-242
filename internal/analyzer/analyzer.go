package analyzer

import (
	"os"
	"path/filepath"

	"code/internal/flags"
)


func isHidden(path string) bool {
	base := filepath.Base(path)
	return len(base) > 0 && string(base[0]) == "."
}

func AnalyzeFile(flags flags.CLIFlags, path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	if !flags.All && isHidden(path) {
		return 0, nil
	}
	return info.Size(), nil

}

func AnalyzeFolder(flags flags.CLIFlags, path string) (int64, error) {
	var totalSize int64 = 0
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	for _, e := range entries {
		fullPath := filepath.Join(path, e.Name())

		// Пропускаем скрытые, если All выключен
		if !flags.All && isHidden(fullPath) {
			continue
		}

		if e.IsDir() {
			if flags.Recursive {
				size, err := AnalyzeFolder(flags, fullPath)
				if err != nil {
					return totalSize, err
				}
				totalSize += size
			}
			// Если не рекурсивно, папки не добавляем
		} else {
			fileInfo, err := e.Info()
			if err != nil {
				return totalSize, err
			}
			totalSize += fileInfo.Size()
		}
	}
	return totalSize, nil
}

// Analyze path size
func Analyze(flags flags.CLIFlags, path string) (int64, error) {

	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return AnalyzeFile(flags, path)
	}
	return AnalyzeFolder(flags, path)
}
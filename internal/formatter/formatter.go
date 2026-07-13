package formatter

import "fmt"

var sizeSuffixes = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

const unitBase int64 = 1024

func formatBytes(size int64) (value float64, unit int) {
	result := float64(size)
	for index := range sizeSuffixes {
		if result >= float64(unitBase) {
			result /= float64(unitBase)
			continue
		}

		return result, index
	}

	return result, len(sizeSuffixes) - 1
}

// humanizeSize size(int64) -> format output(string)
func HumanizeSize(size int64) string {
	if size < unitBase {
		return fmt.Sprintf("%d%s", size, sizeSuffixes[0])
	}

	result, index := formatBytes(size)

	if result == float64(int64(result)) {
		return fmt.Sprintf("%.1f%s", result, sizeSuffixes[index])
	}

	return fmt.Sprintf("%.2f%s", result, sizeSuffixes[index])
}

// GetHumanFormattedResult wrap humanizeSize with cliArgs & path
func GetHumanFormattedResult(human bool, size int64, path string) string {
	var formattedResult string

	if human {
		formattedResult = HumanizeSize(size)
	} else {
		formattedResult = fmt.Sprintf("%d%s", size, sizeSuffixes[0])
	}

	if path != "" {
		return fmt.Sprintf("%s\t%s", formattedResult, path)
	}

	return formattedResult
}

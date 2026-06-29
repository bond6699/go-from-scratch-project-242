package formatter

import "fmt"

var sizeSuffixes = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

const unitBase int64 = 1024

// humanizeSize size(int64) -> format output(string)
func HumanizeSize(size int64) string {
	if size < unitBase {
		return fmt.Sprintf("%d%s", size, sizeSuffixes[0])
	}

	result := float64(size)

	for index := range sizeSuffixes{
		if result >= float64(unitBase) {
			result /= float64(unitBase)
			continue
		}

		if result == float64(int64(result)) {
			return fmt.Sprintf("%.1f%s", result, sizeSuffixes[index])
		}

		return fmt.Sprintf("%.2f%s", result, sizeSuffixes[index])
	}

	if result == float64(int64(result)) {
			return fmt.Sprintf("%.1f%s", result, sizeSuffixes[len(sizeSuffixes)-1])
		}

	return fmt.Sprintf("%.2f%s", result, sizeSuffixes[len(sizeSuffixes)-1])
}

// GetHumanFormattedResult wrap humanizeSize with cliArgs & path
func GetHumanFormattedResult(human bool, size int64, path string) string {
	var formattedResult string

	if human {
		formattedResult = HumanizeSize(size)
	} else {
		formattedResult = fmt.Sprintf("%dB", size)
	}

	if path != "" {
		return fmt.Sprintf("%s\t%s\n", formattedResult, path)
	}

	return formattedResult
}
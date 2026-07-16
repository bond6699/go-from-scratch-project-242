package formatter

import (
	"fmt"
	"math"
)

var sizeSuffixes = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

const unitBase int64 = 1024

func getFormatBytes(size int64) (value float64, suffixIndex int) {
	result := float64(size)
	sizeSuffixesIndex := 0
	for result >= float64(unitBase) {
		sizeSuffixesIndex++
		if result >= float64(unitBase) {
			result /= float64(unitBase)
			continue
		}

		return result, sizeSuffixesIndex
	}

	return result, sizeSuffixesIndex
}

func getHumanSize(human bool, size int64) string {
	if size < unitBase || !human {
		return fmt.Sprintf("%d%s", size, sizeSuffixes[0])
	}

	value, sizeSuffixesIndex := getFormatBytes(size)

	if sizeSuffixesIndex > len(sizeSuffixes)-1 {
		sizeSuffixesIndex = len(sizeSuffixes) - 1
	}

	if math.Mod(value, 1.0) == 0 {
		return fmt.Sprintf("%.1f%s", value, sizeSuffixes[sizeSuffixesIndex])
	}

	return fmt.Sprintf("%.2f%s", value, sizeSuffixes[sizeSuffixesIndex])
}

func GetHumanResult(human bool, size int64, path string) string {
	formattedResult := getHumanSize(human, size)

	if path != "" {
		return fmt.Sprintf("%s\t%s", formattedResult, path)
	}

	return formattedResult
}

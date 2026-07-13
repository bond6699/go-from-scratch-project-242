package formatter

import (
	"fmt"
	"math"
)

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

func HumanizeSize(size int64, human bool) string {
	if size < unitBase || !human {
		return fmt.Sprintf("%d%s", size, sizeSuffixes[0])
	}

	value, unit := formatBytes(size)

	if math.Mod(value, 1.0) == 0 {
		return fmt.Sprintf("%.1f%s", value, sizeSuffixes[unit])
	}

	return fmt.Sprintf("%.2f%s", value, sizeSuffixes[unit])
}

func GetHumanFormattedResult(human bool, size int64, path string) string {
	formattedResult := HumanizeSize(size, human)

	if path != "" {
		return fmt.Sprintf("%s\t%s", formattedResult, path)
	}

	return formattedResult
}

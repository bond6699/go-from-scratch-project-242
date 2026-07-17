package formatter

import (
	"fmt"
	"strconv"
	"strings"
)

var sizeSuffixes = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

const unitBase float64 = 1024

func scaleSize(size int64) (value float64, suffix string) {
	result := float64(size)

	for _, sizeSuffix := range sizeSuffixes {
		if result < unitBase {
			return result, sizeSuffix
		}

		result /= unitBase
	}

	return result, sizeSuffixes[len(sizeSuffixes)-1]
}

func getHumanSize(human bool, size int64) string {
	if size < int64(unitBase) || !human {
		return fmt.Sprintf("%d%s", size, sizeSuffixes[0])
	}

	value, sizeSuffix := scaleSize(size)

	formatValue := strconv.FormatFloat(value, 'f', 2, 64)
	formatValue = strings.TrimSuffix(formatValue, "0")

	return formatValue + sizeSuffix
}

func GetHumanResult(human bool, size int64, path string) string {
	formattedResult := getHumanSize(human, size)

	if path != "" {
		return fmt.Sprintf("%s\t%s", formattedResult, path)
	}

	return formattedResult
}

package formatter

import (
	"fmt"
	"math"
)

func Humanity(size int64) string {
	sizeSuffixes := []string{"Byte", "KB", "MB", "GB"}
	sizeSuffixesIndex := 0
	result := float64(size)
	for result >= float64(1024) {
		result /=  1024
		sizeSuffixesIndex++
	}

	result = math.Floor(result*100)/100

	if result == float64(int64(result)) {
		return fmt.Sprintf("%.0f %s", result, sizeSuffixes[sizeSuffixesIndex])
	}

	return fmt.Sprintf("%.2f %s", result, sizeSuffixes[sizeSuffixesIndex])
}


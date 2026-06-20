package formatter

import (
	"fmt"
	"math"
)

const (
	bytesInKB = 1024
	roundFactor = 100
)

func Humanity(size int64) string {
	sizeSuffixes := []string{"Byte", "KB", "MB", "GB"}
	sizeSuffixesIndex := 0
	result := float64(size)
	for result >= float64(bytesInKB) {
		result /=  bytesInKB
		sizeSuffixesIndex++
	}

	result = math.Floor(result*roundFactor)/roundFactor

	if result == float64(int64(result)) {
		return fmt.Sprintf("%.0f %s", result, sizeSuffixes[sizeSuffixesIndex])
	}

	return fmt.Sprintf("%.2f %s", result, sizeSuffixes[sizeSuffixesIndex])
}


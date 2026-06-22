package formatter

import (
	"code/internal/flags"
	"fmt"
	"math"
)

const (
	bytesInKB   = 1024
	roundFactor = 100
)

func Humanity(size int64) string {
	sizeSuffixes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	sizeSuffixesIndex := 0

	result := float64(size)
	for result >= float64(bytesInKB) {
		result /= bytesInKB
		sizeSuffixesIndex++
	}

	result = math.Floor(result*roundFactor) / roundFactor

	if size < int64(1024) {
		return fmt.Sprintf("%dB", size)
	}

	if result == float64(int64(result)) {
		return fmt.Sprintf("%d.0%s", int64(result), sizeSuffixes[sizeSuffixesIndex])
	}

	return fmt.Sprintf("%.2f%s", result, sizeSuffixes[sizeSuffixesIndex])
}

func Formatter(cmdFlags flags.CLIFlags, size int64) string {
	var result string
	if !cmdFlags.Human {
		result = fmt.Sprintf("%dB", size)
	} else {
		result = Humanity(size)
	}

	return result
}

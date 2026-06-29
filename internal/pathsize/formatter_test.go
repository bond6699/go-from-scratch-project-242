package pathsize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatter(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		expected string
	}{
		{"test1", 25318592, "24.15MB"},
		{"test2", 2077016, "1.98MB"},
		{"test3", 7304971, "6.97MB"},
		{"test4", 7899980, "7.53MB"},
		{"test5", 12745409, "12.15MB"},
		{"test6", 120, "120B"},
		{"test7", 2048, "2.0KB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HumanizeSize(tt.size)
			assert.Equal(t, tt.expected, got)
		})
	}
}

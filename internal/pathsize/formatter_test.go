package pathsize

import "testing"

func TestFormatter(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		expected string
	}{
		{"test1 25318592 -> 24.15MB", 25318592, "24.15MB"},
		{"test2 2077016 -> 1.98MB", 2077016, "1.98MB"},
		{"test3 7304971 -> 6.97MB", 7304971, "6.97MB"},
		{"test4 7899980 -> 7.53MB", 7899980, "7.53MB"},
		{"test5 12745409 -> 12.15MB", 12745409, "12.15MB"},
		{"test6 120 -> 120B", 120, "120B"},
		{"test7 2048 -> 2.00KB", 2048, "2.00KB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := humanizeSize(tt.size)
			if got != tt.expected {
				t.Errorf("got %s, want %s", got, tt.expected)
			}
		})
	}
}

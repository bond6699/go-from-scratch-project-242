package formatter

import "testing"

func TestFormatter(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		expected string
	}{
		{"test1", 25318592, "24.14 MB"},
		{"test2", 2077016, "1.98 MB"},
		{"test3", 7304971, "6.96 MB"},
		{"test4", 7899980, "7.53 MB"},
		{"test5", 12745409, "12.15 MB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Humanity(tt.size)
			if got != tt.expected {
				t.Errorf("got %s, want %s", got, tt.expected)
			}
		})
	}
}

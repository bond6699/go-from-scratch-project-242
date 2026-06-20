package analyzer

import (
	"testing"
	"path/filepath"
    "runtime"

	"code/internal/flags"
)

func getProjectRoot() string {
    _, filename, _, _ := runtime.Caller(0)
    return filepath.Join(filepath.Dir(filename), "../..")
}

func TestAnalyzer(t *testing.T) {
	root := getProjectRoot()
	tests := []struct {
		name string
		flags flags.CLIFlags
		path string
		expected int64
	}{
		{"test1", flags.Create(false, false, false), filepath.Join(root, "tests", "test1"), 25318592},
		{"test2 | Without recursive", flags.Create(false, false, false), filepath.Join(root, "tests", "test2"), 2077016},
		{"test2 | Recursive", flags.Create(true, false, false), filepath.Join(root, "tests", "test2"), 7304971},
		{"test3 | Without recursive, hidden", flags.Create(false, false, false), filepath.Join(root, "tests", "test3"), 7899980},
		{"test3 | Recursive, Hidden", flags.Create(true, true, false), filepath.Join(root, "tests", "test3"), 12745409},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Analyze(tt.flags, tt.path)
			if got != tt.expected {
				t.Errorf("got %d, want %d for path %q with flags %+v", got, tt.expected, tt.path, tt.flags)
			}
		})
	}


}
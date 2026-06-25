package pathsize

import (
	"log"
	"path/filepath"
	"runtime"
	"testing"
)

func getProjectRoot() string {
	_, filename, _, err := runtime.Caller(0)
	if !err {
		log.Fatal("failed to get caller info")
	}

	return filepath.Join(filepath.Dir(filename), "..", "..")
}

func TestAnalyzer(t *testing.T) {
	root := getProjectRoot()
	tests := []struct {
		name     string
		flags    CLIArgs
		path     string
		expected int64
	}{
		{
			"test1 Default ( Recursive false, All false )", Create(false, false, false),
			filepath.Join(root, "tests", "test1"), 25318592,
		},
		{
			"test2 | Folder without recursive ( Recursive false, All false ) ", Create(false, false, false),
			filepath.Join(root, "tests", "test2"), 2077016,
		},
		{
			"test3 | Folder with recursive ( Recursive true, All false )", Create(true, false, false),
			filepath.Join(root, "tests", "test2"), 7304971,
		},
		{
			"test4 | Folder without recursive, with hidden files ( Recursive false, All true )", Create(false, false, false),
			filepath.Join(root, "tests", "test3"), 7899980,
		},
		{
			"test5 | Folder with recursive and hidden files ( Recursive true, All true )", Create(true, true, false),
			filepath.Join(root, "tests", "test3"), 12745409,
		},
		{
			"test6 | File path ( Default => All true, Recursive false )", Create(false, true, false),
			filepath.Join(root, "tests", "test3", "UserGameStats_783965222_220.bin"), 38,
		},
		{
			"test6 | Hidden File path ( Default => All true, Recursive false )", Create(false, true, false),
			filepath.Join(root, "tests", "test3", ".UserGameStatsSchema_550.bin"), 465905,
		},
		
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Analyze(tt.flags, tt.path)
			if got != tt.expected {
				t.Errorf(
					"got %d, want %d for path %q with flags %+v",
					got, tt.expected, tt.path, tt.flags,
				)
			}
		})
	}
}

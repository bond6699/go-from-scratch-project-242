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
			"pathsize/Analyze() test1 Default ( Recursive false, All false )",
			CLIArgs{false, false, false},
			filepath.Join(root, "tests", "test1_base"), 25318592,
		},
		{
			"pathsize/Analyze() test2 | Folder without recursive ( Recursive false, All false ) ",
			CLIArgs{false, false, false},
			filepath.Join(root, "tests", "test2_with_folders"), 2077016,
		},
		{
			"pathsize/Analyze() test3 | Folder with recursive ( Recursive true, All false )",
			CLIArgs{true, false, false},
			filepath.Join(root, "tests", "test2_with_folders"), 7304971,
		},
		{
			"pathsize/Analyze() test4 | Folder without recursive, with hidden files ( Recursive false, All true )",
			CLIArgs{false, true, false},
			filepath.Join(root, "tests", "test3_with_hidden_files_and_folders"), 466268,
		},
		{
			"pathsize/Analyze() test5 | Folder without recursive, without hidden files ( Recursive true, All false )",
			CLIArgs{true, false, false},
			filepath.Join(root, "tests", "test3_with_hidden_files_and_folders"), 439,
		},
		{
			"pathsize/Analyze() test6 | Folder with recursive and hidden files ( Recursive true, All true )",
			CLIArgs{true, true, false},
			filepath.Join(root, "tests", "test3_with_hidden_files_and_folders"), 4845982,
		},
		{
			"pathsize/Analyze() test7 | File path ( Default => All true, Recursive false )",
			CLIArgs{false, true, false},
			filepath.Join(root, "tests", "test3_with_hidden_files_and_folders", "file2.dll"), 38,
		},
		{
			"pathsize/Analyze() test8 | Hidden File path ( Default => All true, Recursive false )",
			CLIArgs{false, true, false},
			filepath.Join(root, "tests", "test3_with_hidden_files_and_folders", ".file1.dll"), 465905,
		},
		{
			"pathsize/Analyze() test9 | Symlink File path ( All false, Recursive false )",
			CLIArgs{false, false, false},
			filepath.Join(root, "tests", "test4_symlink", "file_symlink.dll"), 0,
		},
		{
			"pathsize/Analyze() test10 | Folder with Symlink file ( All false, Recursive false )",
			CLIArgs{false, false, false},
			filepath.Join(root, "tests", "test4_symlink"), 363,
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

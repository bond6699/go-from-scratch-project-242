package pathsize

import (
	"log"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
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
		options  Options
		expected int64
		err      error
	}{
		{
			"pathsize/Analyze() Default",
			Options{
				Recursive: false,
				All:       false,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "base"),
			},
			25318592, nil,
		},
		{
			"pathsize/Analyze() Folder without recursive",
			Options{
				Recursive: false,
				All:       false,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "with_folders"),
			},
			2077016, nil,
		},
		{
			"pathsize/Analyze() Folder with recursive",
			Options{
				Recursive: true,
				All:       false,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "with_folders"),
			},
			7304971, nil,
		},
		{
			"pathsize/Analyze() Folder without recursive, with hidden files",
			Options{
				Recursive: false,
				All:       true,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "with_hidden_files_and_folders"),
			},
			466268, nil,
		},
		{
			"pathsize/Analyze() Folder without recursive, without hidden files",
			Options{
				Recursive: true,
				All:       false,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "with_hidden_files_and_folders"),
			},
			439, nil,
		},
		{
			"pathsize/Analyze() Folder with recursive and hidden files",
			Options{
				Recursive: true,
				All:       true,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "with_hidden_files_and_folders"),
			},
			4845982, nil,
		},
		{
			"pathsize/Analyze() File path",
			Options{
				Recursive: false,
				All:       true,
				Path: filepath.Join(
					root,
					"internal",
					"pathsize",
					"testdata",
					"with_hidden_files_and_folders",
					"file2.dll",
				),
			},
			325,
			nil,
		},
		{
			"pathsize/Analyze() Hidden File path",
			Options{
				Recursive: false,
				All:       true,
				Path: filepath.Join(
					root,
					"internal",
					"pathsize",
					"testdata",
					"with_hidden_files_and_folders",
					".hidden_file.dll",
				),
			},
			465905,
			nil,
		},
		{
			"pathsize/Analyze() Symlink File path",
			Options{
				Recursive: false,
				All:       false,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "symlink", "file_symlink.dll"),
			},
			325, nil,
		},
		{
			"pathsize/Analyze() Folder with Symlink file",
			Options{
				Recursive: false,
				All:       false,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "symlink"),
			},
			688, nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Analyze(tt.options.Recursive, tt.options.All, tt.options.Path)
			assert.Equal(t, tt.expected, got, "path: %s", tt.options.Path)
		})
	}
}

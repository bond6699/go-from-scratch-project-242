package pathsize

import (
	internal_errors "code/internal/errors"
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
			name: "pathsize/Analyze() Default",
			options: Options{
				Recursive: false,
				All:       false,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "base"),
			},
			expected: 25318592, 
			err:      nil,
		},
		{
			name: "pathsize/Analyze() Folder without recursive",
			options: Options{
				Recursive: false,
				All:       false,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "with_folders"),
			},
			expected: 2077016,
			err:      nil,
		},
		{
			name: "pathsize/Analyze() Folder with recursive",
			options: Options{
				Recursive: true,
				All:       false,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "with_folders"),
			},
			expected: 7304971,
			err:      nil,
		},
		{
			name: "pathsize/Analyze() Folder without recursive, with hidden files",
			options: Options{
				Recursive: false,
				All:       true,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "with_hidden_files_and_folders"),
			},
			expected: 466268,
			err:      nil,
		},
		{
			name: "pathsize/Analyze() Folder without recursive, without hidden files",
			options: Options{
				Recursive: true,
				All:       false,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "with_hidden_files_and_folders"),
			},
			expected: 439,
			err:      nil,
		},
		{
			name: "pathsize/Analyze() Folder with recursive and hidden files",
			options: Options{
				Recursive: true,
				All:       true,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "with_hidden_files_and_folders"),
			},
			expected: 4845982,
			err:      nil,
		},
		{
			name: "pathsize/Analyze() File path",
			options: Options{
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
			expected: 325,
			err:      nil,
		},
		{
			name: "pathsize/Analyze() Hidden File path",
			options: Options{
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
			expected: 465905,
			err:      nil,
		},
		{
			name: "pathsize/Analyze() Symlink File path",
			options: Options{
				Recursive: false,
				All:       false,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "symlink", "file_symlink.dll"),
			},
			expected: 325,
			err:      nil,
		},
		{
			name: "pathsize/Analyze() Folder with Symlink file",
			options: Options{
				Recursive: false,
				All:       false,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "symlink"),
			},
			expected: 688,
			err:      nil,
		},
		{
			name: "pathsize/Analyze() Invalid path",
			options: Options{
				Recursive: false,
				All:       false,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "invalid_path"),
			},
			expected: 0,
			err:      internal_errors.ErrReadPath,
		},
		{
			name: "pathsize/Analyze() Access denied",
			options: Options{
				Recursive: false,
				All:       false,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "access_denied"),
			},
			expected: 0,
			err:      internal_errors.ErrAccessDenied,
		},
		{
			name: "pathsize/Analyze() Read path error",
			options: Options{
				Recursive: false,
				All:       false,
				Path:      filepath.Join(root, "internal", "pathsize", "testdata", "read_path_error"),
			},
			expected: 0,
			err:      internal_errors.ErrReadPath,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Analyze(tt.options.Recursive, tt.options.All, tt.options.Path)
			assert.Equal(t, tt.expected, got, "path: %s", tt.options.Path)
			assert.ErrorIs(t, err, tt.err)
		})
	}
}
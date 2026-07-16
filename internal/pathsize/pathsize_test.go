package pathsize

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getProjectRoot() string {
	_, filename, _, err := runtime.Caller(0)
	if !err {
		log.Fatal("failed to get caller info")
	}

	return filepath.Join(filepath.Dir(filename), "..", "..")
}

func TestPermissionDenied(t *testing.T) {
	tempDir := t.TempDir()

	require.NoError(t, os.Chmod(tempDir, 0o000), "не удалось установить права 000")

	t.Cleanup(func() {
		_ = os.Chmod(tempDir, 0o0700)
	})

	result, err := GetPathSize(Options{
		Recursive: false,
		All:       false,
	}, tempDir)
	assert.Equal(t, int64(0), result)
	assert.ErrorIs(t, err, errAccessDenied)
}

func TestAnalyzer(t *testing.T) {
	root := getProjectRoot()
	tests := []struct {
		name        string
		options     Options
		path        string
		expected    int64
		expectedErr error
	}{
		{
			name: "pathsize/Analyze() Default",
			options: Options{
				Recursive: false,
				All:       false,
			},
			path:        filepath.Join(root, "internal", "pathsize", "testdata", "base"),
			expected:    25318592,
			expectedErr: nil,
		},
		{
			name: "pathsize/Analyze() Folder without recursive",
			options: Options{
				Recursive: false,
				All:       false,
			},
			path:        filepath.Join(root, "internal", "pathsize", "testdata", "with_folders"),
			expected:    2077016,
			expectedErr: nil,
		},
		{
			name: "pathsize/Analyze() Folder with recursive",
			options: Options{
				Recursive: true,
				All:       false,
			},
			path:        filepath.Join(root, "internal", "pathsize", "testdata", "with_folders"),
			expected:    7304971,
			expectedErr: nil,
		},
		{
			name: "pathsize/Analyze() Folder without recursive, with hidden files",
			options: Options{
				Recursive: false,
				All:       true,
			},
			path:        filepath.Join(root, "internal", "pathsize", "testdata", "with_hidden_files_and_folders"),
			expected:    466268,
			expectedErr: nil,
		},
		{
			name: "pathsize/Analyze() Folder without recursive, without hidden files",
			options: Options{
				Recursive: true,
				All:       false,
			},
			path:        filepath.Join(root, "internal", "pathsize", "testdata", "with_hidden_files_and_folders"),
			expected:    439,
			expectedErr: nil,
		},
		{
			name: "pathsize/Analyze() Folder with recursive and hidden files",
			options: Options{
				Recursive: true,
				All:       true,
			},
			path:        filepath.Join(root, "internal", "pathsize", "testdata", "with_hidden_files_and_folders"),
			expected:    4845982,
			expectedErr: nil,
		},
		{
			name: "pathsize/Analyze() File path",
			options: Options{
				Recursive: false,
				All:       true,
			},
			path: filepath.Join(
				root,
				"internal",
				"pathsize",
				"testdata",
				"with_hidden_files_and_folders",
				"file2.dll",
			),
			expected:    325,
			expectedErr: nil,
		},
		{
			name: "pathsize/Analyze() Hidden File path",
			options: Options{
				Recursive: false,
				All:       true,
			},
			path: filepath.Join(
				root,
				"internal",
				"pathsize",
				"testdata",
				"with_hidden_files_and_folders",
				".hidden_file.dll",
			),
			expected:    465905,
			expectedErr: nil,
		},
		{
			name: "pathsize/Analyze() Symlink File path",
			options: Options{
				Recursive: false,
				All:       false,
			},
			path:        filepath.Join(root, "internal", "pathsize", "testdata", "symlink", "file_symlink.dll"),
			expected:    325,
			expectedErr: nil,
		},
		{
			name: "pathsize/Analyze() Symlink Folder path",
			options: Options{
				Recursive: false,
				All:       false,
			},
			path:        filepath.Join(root, "internal", "pathsize", "testdata", "folder_symlink"),
			expected:    25318592,
			expectedErr: nil,
		},
		{
			name: "pathsize/Analyze() Folder with Symlink file",
			options: Options{
				Recursive: false,
				All:       false,
			},
			path:        filepath.Join(root, "internal", "pathsize", "testdata", "symlink"),
			expected:    688,
			expectedErr: nil,
		},
		{
			name: "pathsize/Analyze() Invalid path",
			options: Options{
				Recursive: false,
				All:       false,
			},
			path:        filepath.Join(root, "internal", "pathsize", "testdata", "invalid_path"),
			expected:    0,
			expectedErr: errInvalidPath,
		},
		{
			name: "pathsize/Analyze() Empty folder with empty file",
			options: Options{
				Recursive: false,
				All:       false,
			},
			path: filepath.Join(root, "internal", "pathsize", "testdata", "empty_folder"), expected: 0,
			expectedErr: nil,
		},
		{
			name: "pathsize/Analyze() Empty file",
			options: Options{
				Recursive: false,
				All:       false,
			},
			path:        filepath.Join(root, "internal", "pathsize", "testdata", "empty_folder", "empty_file"),
			expected:    0,
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPathSize(tt.options, tt.path)
			assert.Equal(t, tt.expected, got, "path: %s", tt.path)
			assert.ErrorIs(t, err, tt.expectedErr)
		})
	}
}

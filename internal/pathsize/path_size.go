package pathsize

import (
	"os"
	"fmt"
)

type Config struct {
	Path string
	Recursive bool
	Human bool
	All bool
}

func GetFolderSize(path string, recursive, all bool) (int64, error) {
	entities, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var totalSize int64

	for _, entity := range entities {
		fmt.Println("GetFileSize ",	entity.Name())
		if !entity.IsDir() {
			info, err := entity.Info()
			if err != nil {
				return totalSize, err
			}
			fmt.Println("GetFileSize ", info.Name())
			totalSize += info.Size()
		}
	}
	return totalSize, nil
}

func GetFileSize(path string) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return int64(0), err
	}

	fmt.Println("GetFileSize ", info.Name())
	return info.Size(), err
}

func GetPathSize(config Config) (int64, error) {
	info, err := os.Lstat(config.Path)
	if err != nil {
		return int64(0), err
	}

	if info.Mode().Type() == os.ModeSymlink {
		return int64(0), nil
	}

	if !info.IsDir() {
		return GetFileSize(config.Path)
	} else {
		return GetFolderSize(config.Path, config.Recursive, config.All)
	}
}
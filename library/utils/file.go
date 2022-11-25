package utils

import (
	"os"
	"path"
	"strings"
)

func GetAllFilesInDir(dir string) []string {
	result := make([]string, 0, 0)
	entryList, err := os.ReadDir(dir)
	if err != nil {
		return result
	}
	for _, entry := range entryList {
		fullPath := path.Join(dir, entry.Name())
		if entry.IsDir() {
			result = append(result, GetAllFilesInDir(fullPath)...)
		} else {
			result = append(result, fullPath)
		}
	}

	return result
}

func GetFileContent(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	content := string(bytes) // TODO: support more file type
	return content, nil
}

func GetRelativePath(path string, dir string) string {
	prefix := strings.Replace(dir, "./", "", -1)
	index := strings.Index(path, prefix)
	if index == 0 {
		return path[len(prefix)+1:]
	}
	return path
}

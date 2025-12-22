package files

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListFiles(dirPath string) ([]string, error) {
	if err := validateDirExists(dirPath); err != nil {
		return nil, err
	}
	return listWithDepthLimit(dirPath, 3, 0)
}

func validateDirExists(dirPath string) error {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %s", dirPath)
	}
	if !info.IsDir() {
		return fmt.Errorf("path is not a directory: %s", dirPath)
	}
	return nil
}

func listWithDepthLimit(dirPath string, maxDepth int16, currentDepth int16) ([]string, error) {
	var files []string
	if currentDepth > maxDepth {
		return files, nil // Stop recursion if depth limit exceeded
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	for _, entry := range entries {
		fullPath := filepath.Join(dirPath, entry.Name())
		if entry.IsDir() {
			//fmt.Printf("Directory (depth %d): %s\n", currentDepth, fullPath)
			// Recurse into subdirectory
			subDirectoryList, err := listWithDepthLimit(fullPath, maxDepth, currentDepth+1)
			if err != nil {
				return nil, err
			}
			files = append(files, subDirectoryList...)

		} else {
			//fmt.Printf("File (depth %d): %s\n", currentDepth, fullPath)
			files = append(files, fullPath)
		}
	}
	return files, nil
}

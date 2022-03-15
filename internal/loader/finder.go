package loader

import (
	"os"
	"path/filepath"
	"strings"
)

func locateFixtureFile(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return []string{}, err
	}

	return files, nil
}

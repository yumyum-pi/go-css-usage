package html

import (
	"io/fs"
	"path/filepath"
	"strings"
	"yumyum-pi/go-css-usage/pkg/ignorefile"
)

func GetFiles(root string, ignoreList *ignorefile.GitIgnore) ([]string, error) {
	paths := make([]string, 0)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Check if the file matches the ignore pattern.
		//
		if ignoreList.MatchesPath(path) {
			if d.IsDir() {
				return filepath.SkipDir // Skip ignored directories and their contents
			}
			return nil // Skip ignored files
		}

		if !d.IsDir() {
			// check html
			if strings.HasSuffix(path, ".html") {
				paths = append(paths, path)
			}
		}
		return nil
	})
	return paths, err
}

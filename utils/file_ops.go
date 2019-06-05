package utils

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// ReadContents of the given path
func ReadContents(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("os.Open() failed with '%s'\n", err)
	}
	defer f.Close()

	d, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("ioutil.ReadAll() failed with '%s'\n", err)
	}
	return d
}

// FilePathWalkDir - Get all files in directory
func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

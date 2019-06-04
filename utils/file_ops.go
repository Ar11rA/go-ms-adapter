package utils

import (
	"io/ioutil"
	"log"
	"os"
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

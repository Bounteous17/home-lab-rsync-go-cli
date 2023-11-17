package filesystem

import (
	"fmt"
	"io/fs"
	"os"
)

func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func ListFiles(path string) []fs.FileInfo {
	// Open the folder
	dir, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening folder:", err)
		panic("")
	}
	defer dir.Close()

	// Read files in the folder
	fileList, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading folder:", err)
		panic("")
	}

	return fileList
}

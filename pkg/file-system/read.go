package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

type BackupPaths struct {
	Path  string
	IsDir bool
}

var backupPaths []BackupPaths

func visitFolder(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}

	backupPaths = append(backupPaths, BackupPaths{Path: path, IsDir: info.IsDir()})

	return nil
}

func ListFiles(path string) []BackupPaths {
	// Open the folder
	dir, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening folder:", err)
		panic("")
	}
	defer dir.Close()

	// Read files in the folder
	err = filepath.Walk(path, visitFolder)
	if err != nil {
		fmt.Println("Error reading folder:", err)
		panic("")
	}

	return backupPaths
}

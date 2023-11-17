package cmd

import (
	"fmt"

	httpServerSync "github.com/Bounteous17/home-lab-rsync-go-cli/internal/http"
	filesystem "github.com/Bounteous17/home-lab-rsync-go-cli/pkg/file-system"
)

func AppendSyncRequest(backupPath string) {
	// Open the file that you want to send
	backupSubpaths := filesystem.ListFiles(backupPath)

	for _, fileInfo := range backupSubpaths {
		fmt.Println(fileInfo.Path)
		fmt.Printf("fileInfo: %v\n", fileInfo)
		if !fileInfo.IsDir {
			file := filesystem.OpenFile(fileInfo.Path)
			httpServerSync.File(file)
		}
	}
}

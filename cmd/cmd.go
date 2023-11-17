package cmd

import (
	httpServerSync "github.com/Bounteous17/home-lab-rsync-go-cli/internal/http"
	filesystem "github.com/Bounteous17/home-lab-rsync-go-cli/pkg/file-system"
)

func AppendSyncRequest(backupPath string) {
	// Open the file that you want to send
	var fileList = filesystem.ListFiles(backupPath)

	for _, fileInfo := range fileList {
		if !fileInfo.IsDir() {
			file := filesystem.OpenFile(fileInfo.Name())
			httpServerSync.File(file)
		}
	}
}

package main

import (
	"github.com/Bounteous17/home-lab-rsync-go-cli/cmd"
	cliPrompt "github.com/Bounteous17/home-lab-rsync-go-cli/cmd/cli"
)

func main() {
	var backupPath = cliPrompt.AskDesiredBackupPath()
	cmd.AppendSyncRequest(backupPath)
}
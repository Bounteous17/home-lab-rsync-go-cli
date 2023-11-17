package cliPrompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskDesiredBackupPath() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Folder to backup: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	input = strings.TrimSpace(input)
	return input
}

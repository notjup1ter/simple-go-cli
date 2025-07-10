package cmd

import (
	"os"
	"path/filepath"
)

func GetFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "Could not find home directory"
		
	}
	Filepath := filepath.Join(home, "tasks.csv")
	return Filepath
}
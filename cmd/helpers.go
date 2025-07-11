package cmd

import (
	"fmt"
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

func InitializeCSV() {
	file, err := os.OpenFile(GetFilePath(), os.O_RDWR| os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("Error while trying to create file.")
		}
		defer file.Close()
}
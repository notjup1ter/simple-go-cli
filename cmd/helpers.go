package cmd

import (
	"encoding/csv"
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

func getFileAndTasks() (*os.File,[][]string, error) {
	f, err := os.OpenFile(GetFilePath(), os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Could not open file.")
			os.Exit(1)
		}
		
		reader := csv.NewReader(f)
		reader.FieldsPerRecord = 3
		tasks, err := reader.ReadAll()
		if err != nil {
			fmt.Println("Could not read the file")
			os.Exit(1)
		}

		return f, tasks, err
}
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"github.com/olekukonko/tablewriter"
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
	file, err := os.OpenFile(GetFilePath(), os.O_RDWR | os.O_CREATE, 0644)
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
		reader.FieldsPerRecord = 4
		tasks, err := reader.ReadAll()
		if err != nil {
			fmt.Println("Could not read the file")

			os.Exit(1)
		}

		return f, tasks, err
}

func ClearCSV() (*os.File, error) {
	f, err := os.OpenFile(GetFilePath(), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func PrintList() {
	//gathering data
	f, tasks, err := getFileAndTasks()
	if err != nil {
		fmt.Println("Could not get file and its tasks.")
	}
	defer f.Close()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Description", "Status", "Date Added"})
	table.Bulk(tasks)
	table.Render()
}
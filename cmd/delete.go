/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task.",
	Run: func(cmd *cobra.Command, args []string) {
		
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter the id of the task you want to delete: ")
		scanner.Scan()
		id := strings.TrimSpace(scanner.Text())

		f, tasks, err := getFileAndTasks()
		if err != nil { 
			fmt.Println("Could not get the file and its tasks.")
		}

		var updatedTasks [][]string
		for _, task := range tasks {
			if task[0] != id {
				fmt.Println(task[0])
				updatedTasks = append(updatedTasks, task)
			}
		}
		f.Close()

		for i := range updatedTasks {
			updatedTasks[i][0] = strconv.Itoa(i + 1)
		}
		fmt.Println(updatedTasks)

		f, err = ClearCSV()
		if err != nil {
			fmt.Println("Could not clear the CSV file.")
		}

		writer := csv.NewWriter(f)
		err = writer.WriteAll(updatedTasks)
		if err != nil {
			fmt.Println("Could not delete task.")
			return
		}
		f.Close()
		fmt.Println("successfully deleted task")
		PrintList()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

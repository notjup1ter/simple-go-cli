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

// descriptionCmd represents the description command
var descriptionCmd = &cobra.Command{
	Use:   "description",
	Short: "Allows user to update a task's description",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter the id of the task you would like to update: ")
		scanner.Scan()
		trimId := strings.TrimSpace(scanner.Text())
		id,err := strconv.Atoi(trimId)
		if err != nil {
			fmt.Println("Could not convert input to an id.")
		}

		fmt.Println("\n Now, enter the new description")
		scanner.Scan()
		description := strings.TrimSpace(scanner.Text())


		
		f, tasks, err := getFileAndTasks()
		if err != nil {
			fmt.Println("could not get the file and its tasks.")
		}
		defer f.Close()

		for rowIndex := range tasks {
			if rowIndex == id - 1 {
				tasks[rowIndex][1] = description
			} 
		}

		//overwriting the file 
		ClearCSV()

		writer := csv.NewWriter(f)
		err = writer.WriteAll(tasks)
		if err != nil {
			fmt.Println("Could not update task description")
			return
		}
		defer writer.Flush()
		fmt.Println("successfully updated task description")
	},
}

func init() {
	updateTaskCmd.AddCommand(descriptionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// descriptionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// descriptionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

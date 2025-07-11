/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// ncompleteCmd represents the ncomplete command
var ncompleteCmd = &cobra.Command{
	Use:   "ncomplete",
	Short: "Marks a task as incomplete.",
	Long: "",
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//gather user input
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the number of the task you would like to update below:")
		input, err := inputReader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			fmt.Println("Could not read input.")
			return
		}
		inputInt,err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}


		f, err := os.OpenFile(GetFilePath(), os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Could not open file.")
			return
		}
		defer f.Close()

		reader := csv.NewReader(f)
		reader.FieldsPerRecord = 3
		tasks, err := reader.ReadAll()
		if err != nil {
			fmt.Println("Could not read the file")
			return
		}

		for rowIndex := range tasks {
			if rowIndex == inputInt {
				tasks[rowIndex][1] = "not completed"
			}
		}

		//overwriting the file 
		f, err = os.Create(GetFilePath())
		if err != nil {
			fmt.Println("Could not overwrite the file.")
		}
		defer f.Close()

		writer := csv.NewWriter(f)
		err = writer.WriteAll(tasks)
		if err != nil {
			fmt.Println("Could not update task as not completed")
			return
		}
		defer writer.Flush()
		fmt.Println("successfully marked task as not completed")

	},
}

func init() {
	updateTaskCmd.AddCommand(ncompleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ncompleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ncompleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


/*
Copyrig()ht © 2025 NAME HERE <EMAIL ADDRESS>
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

var CFlag bool
var NCFlag bool
// ncompleteCmd represents the ncomplete command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Updates a task's status. -c or -nc required",
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//validate flags
		if CFlag && NCFlag {
			fmt.Println("Cannot use both flags at once")
			os.Exit(1)
		}
		if !CFlag && !NCFlag {
			fmt.Println("Please make sure to use -c for completed or -nc for not completed")
			return
		}


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


		f, tasks := getFileAndTasks()

		for rowIndex := range tasks {
			if rowIndex == inputInt && NCFlag {
				tasks[rowIndex][1] = "not completed"
			} else if rowIndex == inputInt && CFlag {
				tasks[rowIndex][1] = "completed ✅"
			}
		}

		//overwriting the file 
		InitializeCSV()

		writer := csv.NewWriter(f)
		err = writer.WriteAll(tasks)
		if err != nil {
			fmt.Println("Could not update task status")
			return
		}
		defer writer.Flush()
		fmt.Println("successfully updated task's status!")

	},
}

func init() {
	updateTaskCmd.AddCommand(statusCmd)
	statusCmd.Flags().BoolVarP(&CFlag, "completed", "c", false, "marks tasks as completed")
	statusCmd.Flags().BoolVarP(&NCFlag, "ncompleted", "n", false, "marks tasks as Not completed")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ncompleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ncompleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


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
	"time"

	"github.com/spf13/cobra"
)

// addTaskCmd represents the addTask command
var addTaskCmd = &cobra.Command{
	Use:   "addTask",
	Short: "Add a new task to your todo list.",
	Long: `Write a new task to your tasks.csv file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("enter your task below:")
		
		//first, get the task 
		reader := bufio.NewReader(os.Stdin)
		task, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Could not read the input.")
			return
		}
		
		//remove any excess whitespace
		task = strings.TrimSpace(task)

	 	y, m , d := time.Now().Date()
		curr_date :=  strconv.Itoa(y) + " " + m.String() + " " + strconv.Itoa(d)

		taskList := []string{task, "not completed", curr_date}

		//O_append means all new writes are appened at the end, O_WRONLY is write-only" 
		f, err := os.OpenFile(GetFilePath(), os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Could not open file at" + GetFilePath())
			return
		}
		defer f.Close()


		writer := csv.NewWriter(f)

		err = writer.Write(taskList)
		if err != nil { 
			fmt.Println("Could not write the task to the csv file.")
			return
		}
		defer writer.Flush()
		fmt.Println("successfully added task!")
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

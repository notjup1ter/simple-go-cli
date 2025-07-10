/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"encoding/csv"

	"github.com/spf13/cobra"
)

var Fpath string // CL will look for filepath if -f flag is passed

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new CSV file",
	Long: `Creates a CSV file which will contain all of the user's tasks. 
	The csv file will always be stored in the user's home directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		file, err := os.Create(GetFilePath())
		if err != nil {
			fmt.Println("Error while trying to create file.")
		}
		defer file.Close()

		//initalizing the headers
		headers := []string{"Descrition", "Status", "Date Added"}
		writer := csv.NewWriter(file)

		err = writer.Write(headers)
		if err != nil {
			fmt.Println("Could not add the headers to the CSV file.")
		}
		defer writer.Flush()
		fmt.Println("New CSV file has been initalized.")


	},
}


func init() {
	rootCmd.AddCommand(createCmd)
}

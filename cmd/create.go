/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
		
		InitializeCSV()
		fmt.Println("New CSV file has been initalized.")


	},
}


func init() {
	rootCmd.AddCommand(createCmd)
}

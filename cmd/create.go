/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var fpath string // CL will look for filepath if -f flag is passed

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new CSV file",
	Long: `Creates a CSV file which will contain all of the user's tasks. 
	Note, if -f is not included and a filepath is not given, the file will be made in the working directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("filepath") {
			if fpath == "" {
				fmt.Println("Please specify your filepath when using -f")
				return
			}
		}
		fpath = filepath.Join(fpath, "tasks.csv")
		file, err := os.Create(fpath)
		if err != nil {
			fmt.Println("Error while trying to create file.")
		}
		defer file.Close()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	/*StringVarP will assign the arg directly after the flag to the variable specified in the function's parameters,
	while StringP uses the value that is specified directly */

	createCmd.Flags().StringVarP(&fpath, "filepath", "f", "", "Optional: specifies the path of the csv file (do not include the .csv file at the end).")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

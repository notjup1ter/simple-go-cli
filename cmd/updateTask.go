/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// updateTaskCmd represents the updateTask command
var updateTaskCmd = &cobra.Command{
	Use:   "updateTask",
	Short: "Update a task's description or status",
	Long: `Update a task's description or status.
	If the -d (description) flag is present, the selected task's description will be updated instead`,
}

func init() {
	rootCmd.AddCommand(updateTaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateTaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateTaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

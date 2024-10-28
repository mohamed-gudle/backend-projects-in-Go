/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/mohamed-gudle/backend-projects-in-go/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long: `Delete a task from the task tracker.`,
	ValidArgs: []string{"1"},
	Args:  func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires an integer argument")
		}
		if _, err := strconv.Atoi(args[0]); err != nil {
			return fmt.Errorf("invalid integer: %v", args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id,err:=strconv.Atoi(args[0])

		if err!=nil {
			fmt.Errorf(err.Error())
		}


		err=task.DeleteTask(id)

		if err!=nil {
			fmt.Errorf(err.Error())
			
		}

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

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mohamed-gudle/backend-projects-in-go/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

type Task struct {
	ID int
	Description string
	Status string
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a new task",
	Long: `Add a new task to the task tracker.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err:=task.AddTask(args[0])
		if err != nil {
			fmt.Errorf(err.Error())
		}


	

		
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

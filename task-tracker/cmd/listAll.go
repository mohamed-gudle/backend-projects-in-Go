/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/mohamed-gudle/backend-projects-in-go/task-tracker/internal/task"
	"github.com/spf13/cobra"
)

// listAllCmd represents the listAll command
var listAllCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long: `List all tasks in the task tracker.`,
	Run: func(cmd *cobra.Command, args []string) {
		flag:=cmd.Flag("filter")
		filter:=flag.Value.String()

		task.ListTasks(filter)

	},
}

func init() {
	rootCmd.AddCommand(listAllCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listAllCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listAllCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	listAllCmd.Flags().String("filter","all","Choose the status filter: all,done,in-progress")
}

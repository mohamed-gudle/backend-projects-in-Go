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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Status of a task",
	Long: ` Update the status of a task in the task tracker.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id,err:= strconv.Atoi(args[0])
		status:=args[1]

		fmt.Println(id,status,"/n")

		if err!=nil {
			fmt.Errorf(err.Error())
		}

		err=task.UpdateTask(id,status)

		if err!=nil {
			fmt.Errorf(err.Error())
			
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

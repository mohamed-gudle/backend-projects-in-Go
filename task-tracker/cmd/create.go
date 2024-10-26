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
		t:= *task.NewTask(6,args[0],"in-progress")
		err:=task.AddTask(t)
		if err != nil {
			fmt.Errorf(err.Error())
		}


		// file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0755);
		// if err != nil {
		// 	fmt.Println("Error opening file: ", err)
		// 	return
		// }
		// defer file.Close()

		// var tasks []Task
		// data, err:= ioutil.ReadAll(file)
		// if err != nil {
		// 	fmt.Println("Error reading file: ", err)
		// 	return
		// }
		// json.Unmarshal(data, &tasks)
		// tasks = append(tasks, Task{ID: len(tasks) + 1, Description: args[0], Status: "Pending"});
		// data, err = json.Marshal(tasks)

		// if err != nil {
		// 	fmt.Println("Error marshalling data: ", err)
		// 	return
		// }

		// _,err = file.Seek(0, 0);
		// if err != nil {
		// 	fmt.Println("Error seeking file: ", err)
		// 	return
		// }


		// _, err = file.Write(data)
		// if err != nil {
		// 	fmt.Println("Error writing to file: ", err)
		// 	return
		// }


		// defer file.Close()

		
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

package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vpukhanov/task/db"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks the task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil || id <= 0 || id > len(tasks) {
				fmt.Println("invalid task id:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		for _, id := range ids {
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%s\" as completed, error: %s\n", task.Value, err.Error())
			} else {
				fmt.Printf("Marked \"%s\" as completed\n", task.Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}

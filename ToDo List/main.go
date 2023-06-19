package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

type Task struct {
	ID       int
	Title    string
	Due_Date time.Time
	Status   bool
}

var tasks []Task

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "A CLI application to manage tasks",
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get tasks",
	Run: func(cmd *cobra.Command, args []string) {
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		fmt.Println("Tasks:")
		for _, task := range tasks {
			fmt.Printf("ID: %d, Title: %s, Due Date: %s, Status: %t\n", task.ID, task.Title, task.Due_Date, task.Status)
		}
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create task",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		dueDateString, _ := cmd.Flags().GetString("due")
		statusString, _ := cmd.Flags().GetString("status")

		dueDate, err := time.Parse("2006-01-02", dueDateString)

		if err != nil {
			fmt.Println("Invalid date format. Please use yyyy-mm-dd")
		}

		status, err := strconv.ParseBool(statusString)
		if err != nil {
			fmt.Println("Invalid Status value. Please use 'true' or 'false'")
		}

		task := Task{
			ID:       len(tasks) + 1,
			Title:    title,
			Due_Date: dueDate,
			Status:   status,
		}

		tasks = append(tasks, task)

		fmt.Println("Task screated successfully")
	},
}

var UpdateCmd = &cobra.Command{
	Use: "update",
	Short: "Update task",
	Run: func (cmd *cobra.Command, args []string)  {
		id, _ := cmd.Flags().GetString("id")
		
	},
}

func main() {
	

}


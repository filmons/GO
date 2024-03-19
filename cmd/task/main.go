package main

import (
	"fmt"
	"github.com/filmons/semaine2/models" // adjust the import path based on your module name and location
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	err := models.LoadTasks()
	if err != nil {
		log.Fatalf("Failed to load tasks: %v", err)
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List all tasks",
				Action: func(c *cli.Context) error {
					tasks := models.GetAllTasks()
					if len(tasks) == 0 {
						fmt.Println("No tasks found.")
						return nil
					}
					for _, task := range tasks {
						fmt.Printf("ID: %s, Title: %s, Status: %s, Created At: %s\n", task.ID, task.Title, task.Status, task.CreatedAt)
					}
					return nil
				},
			},

			{
				Name:  "show",
				Usage: "Show a task by ID",
				Action: func(c *cli.Context) error {
					id := c.Args().First()
					task, err := models.GetTaskByID(id)
					if err != nil {
						fmt.Println(err)
						return err
					}
					fmt.Printf("%s: %s (Status: %s)\n", task.ID, task.Title, task.Status)
					return nil
				},
			},
			{
				Name:  "add",
				Usage: "Add a new task",
				Action: func(c *cli.Context) error {
					if c.Args().Present() {
						taskTitle := c.Args().First()

						// Assuming AddTask has been updated to modify the global tasks slice and now returns an error if needed
						err := models.AddTask(taskTitle, models.StatusPending) // Ensure AddTask adds the task to your in-memory list and returns any errors
						if err != nil {
							fmt.Println("Error adding task:", err)
							return err
						}

						// After adding a task successfully, save all tasks to the JSON file.
						err = models.SaveTasks(models.GetAllTasks()) // Assuming GetAllTasks returns the current list of tasks
						if err != nil {
							fmt.Println("Failed to save tasks:", err)
							return err
						}

						fmt.Println("Added task:", taskTitle)
					} else {
						fmt.Println("Please provide a task title.")
					}
					return nil
				},
			},

			{
				Name:  "update",
				Usage: "Update an existing task",
				Action: func(c *cli.Context) error {
					if c.Args().Len() < 3 {
						fmt.Println("Usage: update [id] [new title] [new status]")
						return nil
					}
					id := c.Args().Get(0)
					newTitle := c.Args().Get(1)
					newStatus := c.Args().Get(2)

					err := models.UpdateTask(id, newTitle, models.Status(newStatus))
					if err != nil {
						fmt.Printf("Error updating task: %v\n", err)
						return err
					}

					fmt.Println("Task updated successfully")
					return nil
				},
			},

			{
				Name:  "delete",
				Usage: "Delete a task",
				Action: func(c *cli.Context) error {
					if c.Args().Len() < 1 {
						fmt.Println("Usage: delete [id]")
						return nil
					}
					id := c.Args().First()

					err := models.DeleteTask(id)
					if err != nil {
						fmt.Printf("Error deleting task: %v\n", err)
						return err
					}

					fmt.Println("Task deleted successfully")
					return nil
				},
			},
			{
				Name:    "show",
				Usage:   "Show a task by ID",
				Action: func(c *cli.Context) error {
					if c.Args().Len() < 1 {
						fmt.Println("Please provide a task ID.")
						return nil
					}
					id := c.Args().First()
			
					task, err := models.GetTaskByID(id)
					if err != nil {
						fmt.Printf("Error: %v\n", err)
						return err
					}
			
					// Assuming your Task struct includes ID, Title, Status, and CreatedAt fields
					fmt.Printf("ID: %s, Title: %s, Status: %s, Created At: %v\n", task.ID, task.Title, task.Status, task.CreatedAt)
					return nil
				},
			},
			
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli/v3"
)

type Task struct {
	ID          uint64    `json:"id"`
	Description string    `json:"desc"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updateAt,omitempty"`
}

func readTasksFromFile() ([]Task, error) {
	data, err := os.ReadFile("data.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return tasks, nil
}

func writeTasksToFile(tasks []Task) error {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	if err := os.WriteFile("data.json", jsonData, 0o644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

func printTaskInfo(task Task, index int) {
	header := fmt.Sprintf("--------------------- 📚 TASK #%d -----------------------\n", index)
	footer := "--------------------------------------------------------"

	// Create a list of the lines to be printed, using a consistent format
	lines := []string{
		fmt.Sprintf("✨ ID: %d", task.ID),
		fmt.Sprintf("📝 Description: %q", task.Description),
		fmt.Sprintf("🔖 Status: %s", task.Status),
		fmt.Sprintf("📅 Created At: %s", task.CreatedAt.Format("2006-01-02 15:04:05")),
	}

	if !task.UpdatedAt.IsZero() {
		lines = append(lines, fmt.Sprintf("🔄 Updated At: %s", task.UpdatedAt.Format("2006-01-02 15:04:05")))
	}

	// Find the longest line to use for padding calculation
	longestLineLength := len(header) - 1 // -1 to account for the newline
	for _, line := range lines {
		if len(line) > longestLineLength {
			longestLineLength = len(line)
		}
	}

	fmt.Printf("\n%s", header)
	for _, line := range lines {
		// Calculate the padding needed to center the line
		padding := (longestLineLength - len(line)) / 2
		fmt.Printf("%*s%s\n", padding, "", line)
	}
	fmt.Println(footer)
}

func showAllTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for i, task := range tasks {
		printTaskInfo(task, i+1)
	}
}

func showFilteredTasks(tasks []Task, filter string) {
	filteredCount := 0
	for i, task := range tasks {
		if task.Status == filter {
			printTaskInfo(task, i+1)
			filteredCount++
		}
	}
	if filteredCount == 0 {
		fmt.Printf("No tasks found with status: %s\n", filter)
	}
}

func main() {
	cmd := &cli.Command{
		Name:                  "tkm-cli",
		Usage:                 "A simple task-manager made in Go",
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add a new task to your list",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.NArg() < 1 {
						return cli.Exit("Not enough arguments provided", 86)
					}
					tasks, err := readTasksFromFile()

					if err != nil && len(tasks) > 0 {
						log.Fatal(err)
					}
					newID := uint64(len(tasks) + 1)
					newTask := Task{
						ID:          newID,
						Description: cmd.Args().First(),
						Status:      "todo",
						CreatedAt:   time.Now(),
					}
					tasks = append(tasks, newTask)
					if err := writeTasksToFile(tasks); err != nil {
						log.Fatal(err)
					}
					fmt.Printf("Task added successfully! (ID:%d)\n", newTask.ID)
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List all your tasks",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "done",
						Usage: "List done tasks",
					},
					&cli.BoolFlag{
						Name:  "todo",
						Usage: "List todo tasks",
					},
					&cli.BoolFlag{
						Name:  "in-progress",
						Usage: "List in-progress tasks",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					tasks, err := readTasksFromFile()
					if err != nil {
						log.Fatal(err)
					}
					if cmd.Bool("done") {
						showFilteredTasks(tasks, "done")
					} else if cmd.Bool("todo") {
						showFilteredTasks(tasks, "todo")
					} else if cmd.Bool("in-progress") {
						showFilteredTasks(tasks, "in-progress")
					} else {
						showAllTasks(tasks)
					}
					return nil
				},
			},
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "Update a task's description by its ID",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.NArg() < 2 {
						return cli.Exit("Not enough arguments provided. Usage: update <id> <new description>", 86)
					}
					tasks, err := readTasksFromFile()
					if err != nil {
						log.Fatal(err)
					}
					id, err := strconv.ParseUint(cmd.Args().Get(0), 10, 64)
					if err != nil {
						log.Fatalf("Invalid ID format: %v", err)
					}
					updated := false
					for i := range tasks {
						if tasks[i].ID == id {
							tasks[i].Description = cmd.Args().Get(1)
							tasks[i].UpdatedAt = time.Now()
							updated = true
							break
						}
					}
					if !updated {
						fmt.Printf("Task with ID %d not found.\n", id)
						return nil
					}
					if err := writeTasksToFile(tasks); err != nil {
						log.Fatal(err)
					}
					fmt.Printf("Task with ID %d successfully updated.\n", id)
					return nil
				},
			},
			{
				Name:    "mark",
				Aliases: []string{"m"},
				Usage:   "Mark a task as done or in-progress by its ID",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "done",
						Usage: "Mark the task as done",
					},
					&cli.BoolFlag{
						Name:  "in-progress",
						Usage: "Mark the task as in-progress",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.NArg() < 1 {
						return cli.Exit("Task ID is required", 86)
					}
					tasks, err := readTasksFromFile()
					if err != nil {
						log.Fatal(err)
					}
					id, err := strconv.ParseUint(cmd.Args().First(), 10, 64)
					if err != nil {
						log.Fatalf("Invalid ID format: %v", err)
					}
					var newStatus string
					if cmd.Bool("done") {
						newStatus = "done"
					} else if cmd.Bool("in-progress") {
						newStatus = "in-progress"
					} else {
						return cli.Exit("Please specify --done or --in-progress", 86)
					}
					updated := false
					for i := range tasks {
						if tasks[i].ID == id {
							tasks[i].Status = newStatus
							tasks[i].UpdatedAt = time.Now()
							updated = true
							break
						}
					}
					if !updated {
						fmt.Printf("Task with ID %d not found.\n", id)
						return nil
					}
					if err := writeTasksToFile(tasks); err != nil {
						log.Fatal(err)
					}
					fmt.Printf("Task with ID %d successfully marked as %s.\n", id, newStatus)
					return nil
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "Delete a task by its ID",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.NArg() < 1 {
						return cli.Exit("Task ID is required", 86)
					}
					tasks, err := readTasksFromFile()
					if err != nil {
						log.Fatal(err)
					}
					id, err := strconv.ParseUint(cmd.Args().First(), 10, 64)
					if err != nil {
						log.Fatalf("Invalid ID format: %v", err)
					}
					var newTasks []Task
					deleted := false
					for _, task := range tasks {
						if task.ID != id {
							newTasks = append(newTasks, task)
						} else {
							deleted = true
						}
					}
					if !deleted {
						fmt.Printf("Task with ID %d not found.\n", id)
						return nil
					}
					for i := range newTasks {
						newTasks[i].ID = uint64(i + 1)
					}
					if err := writeTasksToFile(newTasks); err != nil {
						log.Fatal(err)
					}
					fmt.Printf("Task with ID %d successfully deleted.\n", id)
					return nil
				},
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

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

type Person struct {
  Id          uint64    `json:"id"`
  Description string    `json:"desc"`
  Status      string    `json:"status"`
  Created_at  time.Time `json:"createdAt"`
  Updated_at  time.Time `json:"updateAt"`
}

func ShowInfo(t []Person){
  c := 1 
  for x := 0; x < len(t); x++{

    fmt.Printf("--------------------- %dº TASK -----------------------\n", c)
    fmt.Printf("Id: %d\nDescription: %q\nStatus: %s\nCreated_at: %s\n", t[x].Id, t[x].Description, t[x].Status, t[x].Created_at.Format("2006-01-02 15:04:05"))
    tk := t[x].Updated_at
    if  tk.String() != "0001-01-01 00:00:00 +0000 UTC"{
      fmt.Printf("Updated_at: %s\n", t[x].Updated_at.Format("2006-01-02 15:04:05"))
    }

    fmt.Println("-------------------------------------------------------")
    c++
  }

}

func ShowInfoWithFilter(t []Person, filter string){
  c := 1 
  for x := 0; x < len(t); x++{

    if t[x].Status == filter {
      fmt.Printf("--------------------- %dº TASK -----------------------\n", c)
      fmt.Printf("Id: %d\nDescription: %q\nStatus: %s\nCreated_at: %s\n", t[x].Id, t[x].Description, t[x].Status, t[x].Created_at.Format("2006-01-02 15:04:05"))
      tk := t[x].Updated_at
      if  tk.String() != "0001-01-01 00:00:00 +0000 UTC"{
        fmt.Printf("Updated_at: %s\n", t[x].Updated_at.Format("2006-01-02 15:04:05"))
      }

      fmt.Println("-------------------------------------------------------")
      c++

    }

  }

}


func MarkTasks( t[]Person, id uint64, filter string){

  for x:= 0; x < len(t); x++ {
    if t[x].Id == uint64(id){
      t[x].Status = filter 
    }
  }

}

func main() {
  cmd := &cli.Command{
    Name:                    "tkm-cli",
    Usage:                   "A simple task-manager made in Go",
    EnableShellCompletion:   true,
    Commands: []*cli.Command{
      {
        Name:    "add",
        Aliases: []string{"a"},
        Usage:   "adding tasks to your list",
        Action: func(ctx context.Context, cmd *cli.Command) error {
          if cmd.NArg() < 1 {
            return cli.Exit("Not enough arguments provided", 86)
          }

          var tasks []Person

          data, err := os.ReadFile("data.json")
          if err != nil {
            if os.IsNotExist(err) {
              tasks = []Person{}
            } else {
              log.Fatal(err)
            }
          } else {
            if err := json.Unmarshal(data, &tasks); err != nil {
              log.Fatal(err)
            }
          }

          var newId uint64 = uint64(len(tasks) + 1)

          newTask := Person{
            Id:          newId,
            Description: cmd.Args().First(),
            Status:      "todo",
            Created_at:  time.Now(),
          }

          tasks = append(tasks, newTask)

          jsonData, err := json.MarshalIndent(tasks, "", "  ")
          if err != nil {
            log.Fatal(err)
          }

          if err := os.WriteFile("data.json", jsonData, 0644); err != nil {
            log.Fatal(err)
          }

          fmt.Printf("Task added successfully! (ID:%d)\n", newTask.Id)
          return nil
        },
      },
      {
        Name:    "list",
        Aliases: []string{"l"},
        Usage:   "list all your tasks",
        Flags: []cli.Flag{
          &cli.BoolFlag{
            Name:  "done",
            Usage: "List the done tasks",
          },
          &cli.BoolFlag{
            Name:  "todo",
            Usage: "List the todo tasks",
          },
          &cli.BoolFlag{
            Name:  "in-progress",
            Usage: "List the in-progress tasks",
          },
        },
        Action: func(ctx context.Context, cmd *cli.Command) error {

          var tasks []Person

          data, err := os.ReadFile("data.json")
          if err != nil {
            if os.IsNotExist(err) {
              tasks = []Person{}
            } else {
              log.Fatal(err)
            }
          } else {
            if err := json.Unmarshal(data, &tasks); err != nil {
              log.Fatal(err)
            }
          }

          if cmd.Bool("done"){
            ShowInfoWithFilter(tasks, "done")
          } else if cmd.Bool("todo"){
            ShowInfoWithFilter(tasks, "todo")
          } else if cmd.Bool("in-progress"){
            ShowInfoWithFilter(tasks, "in-progress")
          } else {
            ShowInfo(tasks)
          }
          return nil
        },
      },
      {
        Name:    "update",
        Aliases: []string{"u"},
        Usage:   "Update your tasks",
        Action: func(ctx context.Context, cmd *cli.Command) error {
          var tasks []Person

          data, err := os.ReadFile("data.json")
          if err != nil {
            if os.IsNotExist(err) {
              tasks = []Person{}
            } else {
              log.Fatal(err)
            }
          } else {
            if err := json.Unmarshal(data, &tasks); err != nil {
              log.Fatal(err)
            }
          }

          i,err := strconv.Atoi(cmd.Args().Get(0))
          if  err != nil{
            log.Fatal(err)
          }

          for x := 0; x < len(tasks); x++{
            if tasks[x].Id == uint64(i){
              tasks[x].Description = cmd.Args().Get(1)
              tasks[x].Updated_at = time.Now() 
            }
          }

          jsonData, err := json.MarshalIndent(tasks, "", "  ")
          if err != nil {
            log.Fatal(err)
          }

          if err := os.WriteFile("data.json", jsonData, 0644); err != nil {
            log.Fatal(err)
          }
          fmt.Printf("The task with id %s has updated\n", cmd.Args().First())

          return nil

        },
      },
      {
        Name:    "mark",
        Aliases: []string{"m"},
        Usage:   "Mark your tasks by category",
        Flags: []cli.Flag{
          &cli.BoolFlag{
            Name:  "done",
            Usage: "Mark the done task",
          },
          &cli.BoolFlag{
            Name:  "in-progress",
            Usage: "Mark the in-progress task",
          },
        },
        Action: func(ctx context.Context, cmd *cli.Command) error {

          var tasks []Person

          data, err := os.ReadFile("data.json")
          if err != nil {
            if os.IsNotExist(err) {
              tasks = []Person{}
            } else {
              log.Fatal(err)
            }
          } else {
            if err := json.Unmarshal(data, &tasks); err != nil {
              log.Fatal(err)
            }
          }

          i, err := strconv.Atoi(cmd.Args().First())
          if err != nil {
            log.Fatal(err)
          }

          if cmd.Bool("done"){
            MarkTasks(tasks, uint64(i), "done")
          } else if cmd.Bool("in-progress"){
            MarkTasks(tasks, uint64(i), "in-progress")
          }

          jsonData, err := json.MarshalIndent(tasks, "", "  ")
          if err != nil {
            log.Fatal(err)
          }

          if err := os.WriteFile("data.json", jsonData, 0644); err != nil {
            log.Fatal(err)
          }

          fmt.Printf("The task ID - %s sucessfully marked\n", cmd.Args().First())

          return nil
        },
      },
      {
        Name:    "delete",
        Aliases: []string{"d"},
        Usage:   "Delete your tasks",
        Action: func(ctx context.Context, cmd *cli.Command) error {
          var tasks []Person

          var newTasks []Person

          data, err := os.ReadFile("data.json")
          if err != nil {
            if os.IsNotExist(err) {
              tasks = []Person{}
            } else {
              log.Fatal(err)
            }
          } else {
            if err := json.Unmarshal(data, &tasks); err != nil {
              log.Fatal(err)
            }
          }

          i, err := strconv.Atoi(cmd.Args().First())
            newId := uint64(1)
          for x := 0; x < len(tasks); x++{

            if tasks[x].Id != uint64(i){
              tasks[x].Id = newId
              newTasks = append(newTasks, tasks[x])
              newId++
            } 
          }
          jsonData, err := json.MarshalIndent(newTasks, "", "  ")
          if err != nil {
            log.Fatal(err)
          }

          if err := os.WriteFile("data.json", jsonData, 0644); err != nil {
            log.Fatal(err)
          }

          fmt.Printf("The task ID - %s sucessfully deleted\n", cmd.Args().First())

          return nil
        },
      },
    },
  }

  if err := cmd.Run(context.Background(), os.Args); err != nil {
    log.Fatal(err)
  }
}


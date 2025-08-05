package main

import (
  "os"
  _ "fmt"
  "context"
  "log"
  "github.com/urfave/cli/v3"
  _ "encoding/json"
  _ "time"
)

type Person struct {
  Id uint64 `json:"id"`
  Description string `json:"desc"`
  Status string `json:"status"`
  Created_at *cli.TimestampFlag `json:"createdAt"`
  Updated_at *cli.TimestampFlag `json:"updateAt"`
}

var person Person

func main (){
  cmd := &cli.Command{
    Name : "tkm-cli",
    Usage: "A simple task-manager made in Go :)",
    EnableShellCompletion: true,
    Commands: []*cli.Command{
      {
        Name : "add",
        Aliases : []string{"a"},
        Usage : "adding tasks to your list",
        Action : func (ctx context.Context, cmd *cli.Command) error {

          return nil
        },
      },
      {
        Name : "list",
        Aliases : []string{"l"},
        Usage : "list all your tasks ",
        Flags: []cli.Flag{
          &cli.BoolFlag{
              Name : "done",
              Usage : "List the done tasks",
          },
          &cli.BoolFlag{
              Name : "todo",
              Usage : "List the todo tasks",
          },
          &cli.BoolFlag{
              Name : "in-progress",
              Usage : "List the in-progress tasks",
          },
        },
        Action: func (ctx context.Context, cmd *cli.Command) error {
          return nil
        },
      },
      {
        Name : "update",
        Aliases : []string{"u"},
        Usage : "Update your tasks",
        Action : func (ctx context.Context, cmd *cli.Command) error {

          return nil

        },
      },
      {
        Name : "mark",
        Aliases : []string{"m"},
        Usage : "Mark your tasks by categoria",
        Flags: []cli.Flag{
        &cli.Uint64Flag{
          Name : "done",
          Usage : "Mark the done task",
        },
        &cli.Uint64Flag{
          Name : "in-progress",
          Usage : "Mark the in-progress task",
        },
        },
        Action : func (ctx context.Context, cmd *cli.Command) error {

          return nil

        },
      },
      {
        Name : "delete",
        Aliases : []string{"d"},
        Usage : "Delete your tasks",
        Action : func (ctx context.Context, cmd *cli.Command) error {

          return nil

        },
      },
    },
  }

  if err := cmd.Run(context.Background(), os.Args); err != nil {
    log.Fatal(err)
  }
}

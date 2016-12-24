package main

import (
  "os"
  "fmt"

  "github.com/urfave/cli"
)

func main()  {

  app := cli.NewApp()
  app.Name = "M&M"
  app.Usage = "one day we will build an actually impressive Go script"
  app.Action = func (c *cli.Context) error {
    fmt.Println("Hi! Welcome to the M&M commandline app\n")
    fmt.Println("It was created by Max and Michael.\nWe couldn't really decide who's name goes first in the title, hence M&M")

    return nil
  }

  app.Run(os.Args)
}

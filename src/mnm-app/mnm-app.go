package main

import (
  "os"
  "fmt"

  "github.com/urfave/cli"
)

func main()  {

  app := cli.NewApp()
  app.Name = "boom"
  app.Usage = "make an explosive entrance"
  app.Action = func (c *cli.Context) error {
    fmt.Println("boom! I say!")
    return nil
  }

  app.Run(os.Args)
}

package main

import (
  "os"
  "fmt"

  "github.com/urfave/cli"
)

func main()  {
  var genre string

  app := cli.NewApp()

  app.Name = "M&M"
  app.Usage = "one day we will build an actually impressive Go script"

  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "genre, g",
      Value: "future bass",
      Usage: "deciding what music to offer",
      Destination: &genre,
    },
  }

  app.Action = func (c *cli.Context) error {
    fmt.Println("\nHi! Welcome to the M&M commandline app\n")
    fmt.Println("This is not your simple cli app - it can understand flags, too.\nBased on user input, the app returns a song of a certain genre.\n")
    fmt.Printf("Listen to ")

    if genre == "rap" {
      fmt.Printf("Kanye West - New Slaves")
    } else {
      fmt.Printf("Marhsmello - Alone")
    }

    return nil
  }

  app.Run(os.Args)
}

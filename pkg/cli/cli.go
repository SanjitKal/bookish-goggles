package cli

import (
	"github.com/urfave/cli"
	"os"
)

func StartCli() {
  app := cli.NewApp()
  app.Commands = []cli.Command{
    {
      Name:    "get",
      Aliases: []string{"g"},
      Usage:   `"get k" will retrieve the value of key 'k' from the store. If key 'k' does not exist in the store, "NOT_FOUND" will be displayed.`,
      Action:  func(c *cli.Context) error {
        return nil
      },
    },
    {
      Name:    "put",
      Aliases: []string{"p"},
      Usage:   `"put k v" will insert the key 'k' in the store, bound to the value 'v'. If key 'k' already exists in the store, its' value will be updated. Values are assumed to be quoted strings.`,
      Action:  func(c *cli.Context) error {
        return nil
      },
    },
  }

  sort.Sort(cli.FlagsByName(app.Flags))
  sort.Sort(cli.CommandsByName(app.Commands))

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
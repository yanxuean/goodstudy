package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	var language string
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang",
			Value:       "English",
			Usage:       "language for the greeting",
			Destination: &language,
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Printf("boom! I say! %q\n", c.Args().Get(0))
		name := "someone"
		if c.NArg() > 0 {
			name = c.Args()[0]
		}
		if language == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
		return nil
	}
	app.Run(os.Args)

	//fmt.Println("vim-go")
}

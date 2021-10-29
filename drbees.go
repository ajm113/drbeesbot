package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "drbees"
	app.Usage = "DR BEES Twitter bot using Golang that's not devoided of any bees"
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			TakesFile: true,
			Name:      "config",
			Aliases:   []string{"c"},
			Value:     "config.yml",
			Usage:     "configuration file used to connect to twitter and change beehavior.",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Action:  run,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

package cli

import (
	"github.com/urfave/cli/v2"
)

func GetCli() *cli.App {
	app := &cli.App{
		Name:      "chk",
		Usage:     "A simple command line tool to check connection between devices.",
		UsageText: "chk [command flags]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "type",
				Aliases: []string{"t"},
				Usage:   "Determines the type of connection that will be tested.",
			},
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}
	return app
}

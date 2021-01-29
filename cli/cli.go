package cli

import (
	"github.com/Ripolak/chk/chk"
	"github.com/urfave/cli/v2"
)

func GetCli() *cli.App {
	app := &cli.App{
		Name:      "chk",
		Usage:     "A simple command line tool to check connection between devices.",
		UsageText: "chk [command flags]",
		Commands: []*cli.Command{
			{
				Name:    "ping",
				Aliases: []string{"icmp"},
				Usage:   "Check if the target address is accessible via ICMP protocol.",
				Action: func(c *cli.Context) error {
					return chk.CheckICMP()
				},
			},
		},
	}
	return app
}

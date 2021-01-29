package cli

import (
	"fmt"

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
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "count",
						Aliases: []string{"c"},
						Value:   1,
						Usage:   "Amount of ping requests that will be sent. Defaults to 1.",
					},
					&cli.IntFlag{
						Name:    "timeout",
						Aliases: []string{"t"},
						Value:   1,
						Usage:   "Time to wait for a response in seconds. Defaults to 1.",
					},
					&cli.IntFlag{
						Name:    "interval",
						Aliases: []string{"i"},
						Value:   1,
						Usage:   "Time to wait between ping requests in seconds if the count is more than 1. Defaults to 1.",
					},
				},
				Action: func(c *cli.Context) error {
					address := c.Args().Get(0)
					count := c.Int("count")
					timeout := c.Int("timeout")
					interval := c.Int("interval")
					res, err := chk.CheckICMP(address, count, timeout, interval)
					fmt.Print(res.PacketLoss)
					return err
				},
			},
		},
	}
	return app
}

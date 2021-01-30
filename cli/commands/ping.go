package commands

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Ripolak/chk/chk"
	"github.com/go-ping/ping"
	"github.com/urfave/cli/v2"
)

// PingCommand represents the 'ping' or 'icmp' commnand options for the CLI.
var PingCommand = &cli.Command{
	Name:    "ping",
	Aliases: []string{"icmp"},
	Usage:   "Check if the target address is accessible via ICMP protocol.",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "count",
			Aliases: []string{"c"},
			Value:   1,
			Usage:   "Amount of ping requests that will be sent.",
		},
		&cli.IntFlag{
			Name:    "timeout",
			Aliases: []string{"t"},
			Value:   1,
			Usage:   "Time to wait for a response in seconds.",
		},
		&cli.IntFlag{
			Name:    "interval",
			Aliases: []string{"i"},
			Value:   1,
			Usage:   "Time to wait between ping requests in seconds if the count is more than 1.",
		},
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Value:   false,
			Usage:   "Specifies whether to show verbose information about the ping result.",
		},
	},
	Action: pingAction,
}

func pingAction(c *cli.Context) error {
	address := c.Args().First()
	err := validateArgument(address)
	if err != nil {
		return err
	}
	verbose := c.Bool("verbose")
	count := c.Int("count")
	timeout := c.Int("timeout")
	interval := c.Int("interval")
	res, err := chk.CheckICMP(address, count, timeout, interval)
	displayPingResult(res, address, verbose)
	return nil
}

func displayPingResult(res *ping.Statistics, address string, verbose bool) {
	if verbose {
		b, err := json.MarshalIndent(res, "", "	")
		if err != nil {
			fmt.Println("Failed to parse verbose info about the ping result.")
		} else {
			fmt.Println(string(b))
		}
	}
	if !reflect.ValueOf(res).IsNil() && res.PacketLoss != 100 {
		fmt.Println(fmt.Sprintf("Successful ping connection to %s", address))
	} else {
		fmt.Println(fmt.Sprintf("%s cannot be reached.", address))
	}
}

func validateArgument(address string) error {
	if address == "" {
		return fmt.Errorf("address must be provided")
	}
	return nil
}

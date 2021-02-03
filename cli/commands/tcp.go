package commands

import (
	"fmt"
	"net"

	"github.com/Ripolak/chk/connection"
	"github.com/urfave/cli/v2"
)

// TCPCommand represents a command in the CLI that checks TCP connection.
var TCPCommand = &cli.Command{
	Name:  "tcp",
	Usage: "Check if the target address is accessible via TCP.",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "timeout",
			Aliases: []string{"t"},
			Value:   1,
			Usage:   "Time to wait for a response in seconds.",
		},
	},
	Action: tcpAction,
}

func tcpAction(c *cli.Context) error {
	address := c.Args().First()
	err := validateArgument(address)
	timeout := c.Int("timeout")
	if err != nil {
		return err
	}
	res, err := connection.CheckTCP(address, timeout)
	if err != nil {
		displayTCPErrorResult(err, address)
	} else {
		displayTCPResult(res, address)
	}
	return nil
}

func displayTCPResult(res net.Conn, address string) {
	fmt.Println(fmt.Sprintf("Successful TCP connection to %s.", address))
}

func displayTCPErrorResult(err error, address string) {
	fmt.Println(fmt.Sprintf("TCP Connection to %s failed. Error that was received: '%s'", address, err.Error()))
}

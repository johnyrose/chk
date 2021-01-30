package commands

import (
	"fmt"
	"net"

	"github.com/Ripolak/chk/chk"
	"github.com/urfave/cli/v2"
)

// TCPCommand represents a command in the CLI that checks TCP connection.
var TCPCommand = &cli.Command{
	Name:   "tcp",
	Usage:  "Check if the target address is accessible via TCP.",
	Action: tcpAction,
}

func tcpAction(c *cli.Context) error {
	address := c.Args().First()
	err := validateArgument(address)
	if err != nil {
		return err
	}
	res, err := chk.CheckTCP(address)
	if err != nil {
		displayTCPResult(res, address)
	} else {
		displayTCPErrorResult(err, address)
	}
	return nil
}

func displayTCPResult(res *net.TCPConn, address string) {
	fmt.Println(fmt.Sprintf("Successful TCP connection to %s.", address))
}

func displayTCPErrorResult(err error, address string) {
	fmt.Println(fmt.Sprintf("TCP Connection to %s failed. Error that was received: '%s'", address, err.Error()))
}

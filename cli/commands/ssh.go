package commands

import (
	"fmt"
	"net"
	"strings"

	"github.com/Ripolak/chk/chk"
	"github.com/urfave/cli/v2"
)

// SSHCommand represents a command in the CLI that checks SSH connection.
var SSHCommand = &cli.Command{
	Name:  "ssh",
	Usage: "Check if the target address is accessible via SSH.",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "timeout",
			Aliases: []string{"t"},
			Value:   1,
			Usage:   "Time to wait for a response in seconds.",
		},
	},
	Action: sshAction,
}

func sshAction(c *cli.Context) error {
	address := c.Args().First()
	err := validateArgument(address)
	timeout := c.Int("timeout")
	if err != nil {
		return err
	}
	_, err = chk.CheckSSH(address, timeout)
	displaySSHResult(address, err)
	return nil
}

func displaySSHResult(address string, err error) {
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			fmt.Println(fmt.Sprintf("SSH connection to %s failed. The following error was received: %s", address, err.Error()))
		} else if strings.Contains(err.Error(), "overflow reading") {
			fmt.Println(fmt.Sprintf("SSH connection to %s failed. Target is reachable but doesn't appear to listen in SSH protocol or uses a non-mathching SSH version.", address))
		} else if strings.Contains(err.Error(), "unable to authenticate") {
			fmt.Println(fmt.Sprintf("Successful SSH connection to %s.", address))
		}
	} else {
		fmt.Println(fmt.Sprintf("Successful SSH connection to %s.", address))
	}
}

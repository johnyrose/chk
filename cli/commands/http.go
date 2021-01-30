package commands

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Ripolak/chk/chk"
	"github.com/urfave/cli/v2"
)

// HttpCommand represents a command in the CLI that checks HTTP or HTTPS connection.
var HttpCommand = &cli.Command{
	Name:  "http",
	Usage: "Check if the target address is accessible via HTTP.",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "timeout",
			Aliases: []string{"t"},
			Value:   1,
			Usage:   "Time to wait for a response in seconds.",
		},
	},
	Action: httpAction,
}

func httpAction(c *cli.Context) error {
	address := c.Args().First()
	err := validateArgument(address)
	if err != nil {
		return err
	}
	timeout := c.Int("timeout")
	res, err := chk.CheckHTTP(address, timeout)
	if err != nil {
		displayHttpErrorResult(address, err)
	} else {
		displayHttpResult(res.Response(), address)
	}
	return nil
}

func displayHttpResult(res *http.Response, address string) {
	fmt.Println(fmt.Sprintf("Successful http connection to %s. Response code that was received: %v", address, res.StatusCode))
}

func displayHttpErrorResult(address string, err error) {
	if strings.Contains(err.Error(), "connection refused") {
		fmt.Println(fmt.Sprintf("Got connection refused from %s. Target is reachable but does not listen in the specified port.", address))
	} else {
		fmt.Println(fmt.Sprintf("HTTP connection to %s failed. The following error was received: '%s'", address, err))
	}
}

package commands

import (
	"github.com/Ripolak/chk/chk"
	"github.com/imroc/req"
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
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Value:   false,
			Usage:   "Specifies whether to show verbose information about the http result.",
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
	verbose := c.Bool("verbose")
	timeout := c.Int("timeout")
	res, err := chk.CheckHTTP(address, timeout)
	displayHttpResult(res, address, verbose)
	return nil
}

func displayHttpResult(res *req.Resp, address string, verbose bool) {

}

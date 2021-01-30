package cli

import (
	"github.com/Ripolak/chk/cli/commands"
	"github.com/urfave/cli/v2"
)

func GetCli() *cli.App {
	app := &cli.App{
		Name:      "chk",
		Usage:     "A simple command line tool to check connection between devices.",
		UsageText: "chk [command flags]",
		Commands: []*cli.Command{
			commands.PingCommand,
			commands.HttpCommand,
			commands.TCPCommand,
		},
	}
	return app
}

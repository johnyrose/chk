package cli

import (
	"github.com/urfave/cli"
)

func GetCli() *cli.App {
	app := &cli.App{
		Name:      "chk",
		Usage:     "A simple command line tool to check connection between devices.",
		UsageText: "chk [command flags]",
		Commands:  []cli.Command{},
	}
	return app
}

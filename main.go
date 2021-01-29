package main

import (
	"fmt"
	"os"

	"github.com/Ripolak/chk/cli"
)

func main() {
	app := cli.GetCli()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

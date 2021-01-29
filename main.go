package main

import (
	"log"
	"os"

	"github.com/Ripolak/chk/cli"
)

func main() {
	app := cli.GetCli()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

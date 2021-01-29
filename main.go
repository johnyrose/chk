package main

import (
	"chk/cli"
	"log"
	"os"
)

func main() {
	app := cli.GetCli()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"os"

	"gocar/internal/cli"
)

func main() {
	app := cli.NewApp()
	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

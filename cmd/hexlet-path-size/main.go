package main

import (
	"code/internal/cli"
	"context"
	"fmt"
	"os"
)

func main() {
	app := cli.CreateApp()

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

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
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}

	os.Exit(0)
}

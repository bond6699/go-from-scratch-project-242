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
		fmt.Fprintf(os.Stderr, "%v\n", err)
		fmt.Fprintf(os.Stdout, "Try use --help\n")

		os.Exit(1)
	}

	os.Exit(0)
}

package main

import (
	"code/internal/cli"
	"fmt"
	"os"
)

func main() {
	err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		fmt.Fprintf(os.Stderr, "Try use --help\n")

		os.Exit(1)
	}

	os.Exit(0)
}

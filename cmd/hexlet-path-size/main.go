package main

import (
    "os"
    "context"
	"fmt"
	"log"
    "github.com/urfave/cli/v3"
)

func Action(ctx context.Context, cmd *cli.Command) error {
	fmt.Println("Hello World")
	return nil
}

func main() {
	cmd := &cli.Command{
		Name: "hexlet-path-size",
		Usage: "print size of a file or directory",
		UsageText: "hexlet-path-size [global options] <path>",
		Action: Action,
	}

    if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
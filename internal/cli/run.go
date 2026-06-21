// Package cli provides functions for run cli util.
package cli

import (
	"code/internal/analyzer"
	"code/internal/flags"
	"code/internal/formatter"
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

// Error: path required.
var ErrPathRequired = errors.New("path is required")

// ActionLogic processes the CLI command and prints the file/directory size.
//
//nolint:forbidigo
func ActionLogic(ctx context.Context, cmd *cli.Command) error {
	if cmd.NArg() == 0 {
		return ErrPathRequired
	}

	path := cmd.Args().First()
	cliflags := flags.Create(
		cmd.Bool("recursive"),
		cmd.Bool("all"),
		cmd.Bool("human"),
	)

	result, err := analyzer.Analyze(cliflags, path)
	if err != nil {
		return fmt.Errorf("analyzing path %s: %w", path, err)
	}

	formattedResult := formatter.Formatter(cliflags, result)

	fmt.Println(formattedResult)

	return nil
}

// Run CLI Util.
//
//nolint:exhaustruct
func Run() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory",
		UsageText: "hexlet-path-size [global options] <path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "recursive", Aliases: []string{"r"},
				Usage: "recursive size of directories.",
			},
			&cli.BoolFlag{
				Name: "all", Aliases: []string{"a"},
				Usage: "include hidden files and directories.",
			},
			&cli.BoolFlag{
				Name: "human", Aliases: []string{"H"},
				Usage: "human-readable sizes (auto-select unit)",
			},
		},
		Action: ActionLogic,
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

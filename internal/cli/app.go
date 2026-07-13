// Package cli provides functions for run cli util.
package cli

import (
	"code/internal/formatter"
	"code/internal/pathsize"
	"context"
	"errors"
	"fmt"

	"github.com/urfave/cli/v3"
)

var errPathArgsNotProvided = errors.New("path argument not provided, only one path argument is allowed")

func parseOptions(cmd *cli.Command) (pathsize.Options, error) {
	if cmd.NArg() != 1 {
		return pathsize.Options{}, errPathArgsNotProvided
	}

	return pathsize.Options{
		Recursive: cmd.Bool("recursive"),
		All:       cmd.Bool("all"),
		Path:      cmd.Args().First(),
	}, nil
}

func actionLogic(ctx context.Context, cmd *cli.Command) error {
	options, err := parseOptions(cmd)
	if err != nil {
		return err
	}

	result, err := pathsize.Analyze(options.Recursive, options.All, options.Path)
	if err != nil {
		return err
	}

	fmt.Println(formatter.GetHumanFormattedResult(
		cmd.Bool("human"),
		result,
		options.Path,
	))

	return nil
}

func CreateApp() *cli.Command {
	app := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		UsageText: "hexlet-path-size [global options] <path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "recursive", Aliases: []string{"r"},
				Usage: "recursive size of directories (default: false)",
			},
			&cli.BoolFlag{
				Name: "human", Aliases: []string{"H"},
				Usage: "human-readable sizes (auto-select unit) (default: false)",
			},
			&cli.BoolFlag{
				Name: "all", Aliases: []string{"a"},
				Usage: "include hidden files and directories (default: false)",
			},
		},
		Action: actionLogic,
	}

	return app
}

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

var errPathArgsCount = errors.New("invalid argument count")

func parseOptions(cmd *cli.Command) (pathsize.Options, error) {
	path := cmd.StringArg("path")

	if path == "" {
		return pathsize.Options{}, errPathArgsCount
	}

	return pathsize.Options{
		Recursive: cmd.Bool("recursive"),
		All:       cmd.Bool("all"),
		Path:      path,
	}, nil
}

func actionLogic(ctx context.Context, cmd *cli.Command) error {
	options, err := parseOptions(cmd)
	if err != nil {
		_ = cli.ShowRootCommandHelp(cmd)
		return err
	}

	result, err := pathsize.Analyze(options)
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
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:      "path",
				UsageText: "<path>",
			},
		},
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

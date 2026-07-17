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

func getPath(cmd *cli.Command) (string, error) {
	path := cmd.StringArg("path")

	if path == "" {
		return "", errPathArgsCount
	}

	return path, nil
}

func getOptions(cmd *cli.Command) pathsize.Options {
	return pathsize.Options{
		Recursive: cmd.Bool("recursive"),
		All:       cmd.Bool("all"),
	}
}

func actionLogic(ctx context.Context, cmd *cli.Command) error {
	path, err := getPath(cmd)
	if err != nil {
		_ = cli.ShowRootCommandHelp(cmd)
		return err
	}

	options := getOptions(cmd)

	result, err := pathsize.GetPathSize(options, path)
	if err != nil {
		return err
	}

	fmt.Println(formatter.GetHumanResult(
		cmd.Bool("human"),
		result,
		path,
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

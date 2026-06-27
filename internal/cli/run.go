// Package cli provides functions for run cli util.
package cli

import (
	"code/internal/pathsize"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v3"
)

// cliArgsParse parse CLI flags ,path and validate path
func cliArgsParse(cmd *cli.Command) (pathsize.CLIArgs, string, error) {
	args := cmd.Args().Slice()
	if len(args) > 1 {
		return pathsize.CLIArgs{}, "", fmt.Errorf(
			"error: expected 1 path, got %d. Usage: hexlet-path-size [options] <path>",
			len(args),
		)
	} else if len(args) == 0 {
		return pathsize.CLIArgs{}, "", fmt.Errorf(
			"error: expected 1 path, got 0. Usage: hexlet-path-size [options] <path>",
		)
	}

	path := filepath.Clean(args[0])

	_, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return pathsize.CLIArgs{}, "", fmt.Errorf("error: Received path \"%s\" is not exist", path)
		}

		return pathsize.CLIArgs{}, "", fmt.Errorf("error with received path \"%s\"", path)
	}

	if pathsize.IsSymLink(path) {
		return pathsize.CLIArgs{}, "", fmt.Errorf("error: Received path \"%s\" is sumlink", path)
	}

	return pathsize.CLIArgs{
		Recursive: cmd.Bool("recursive"),
		Human:     cmd.Bool("human"),
		All:       cmd.Bool("all"),
	}, path, nil
}

func actionLogic(ctx context.Context, cmd *cli.Command) error {
	cliArgs, path, err := cliArgsParse(cmd)
	if err != nil {
		return err
	}

	result, err := pathsize.Analyze(cliArgs, path)
	if err != nil {
		return fmt.Errorf("error analyzing path %s: %w", path, err)
	}

	formattedResult := pathsize.GetHumanFormattedResult(cliArgs, result, path)

	fmt.Println(formattedResult)

	return nil
}

// Run CLI Util.
func Run() error {
	cmd := &cli.Command{
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

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		return err
	}

	return nil
}

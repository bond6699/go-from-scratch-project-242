// Package cli provides functions for run cli util.
package cli

import (
	"code/internal/formatter"
	"code/internal/pathsize"
	internal_errors "code/internal/errors"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"github.com/urfave/cli/v3"
)

func validatePath(path string) error {
	_, err := os.Lstat(path)
	if err != nil {
		return fmt.Errorf("path %q: %w", path, internal_errors.ErrInvalidPath)
	}
	return nil
}

// cliArgsParse parse CLI flags ,path and validate path
func parseOptions(cmd *cli.Command) (pathsize.Options, error) {
	if cmd.NArg() != 1 {
		return pathsize.Options{}, internal_errors.ErrPathArgsNotProvided
	}

	path := filepath.Clean(cmd.Args().Get(0))
	err := validatePath(path)
	if err != nil {
		return pathsize.Options{}, err
	}

	return pathsize.Options{
		Recursive: cmd.Bool("recursive"),
		Human:     cmd.Bool("human"),
		All:       cmd.Bool("all"),
		Path:      path,
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
		options.Human, 
		result, 
		options.Path,
	)) 

	return nil
}

// Run CLI Util.
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
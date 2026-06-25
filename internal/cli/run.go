// Package cli provides functions for run cli util.
package cli

import (
	"code/internal/pathsize"
	"path/filepath"
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func cliArgsParse(cmd *cli.Command) (pathsize.CLIArgs, string, error) {
	args := cmd.Args().Slice()
	if len(args) > 1 {
		return pathsize.CLIArgs{}, "", fmt.Errorf("Error: expected 1 path, got %d. Usage: hexlet-path-size [options] <path>", len(args))
	} else if len(args) == 0 {
		return pathsize.CLIArgs{}, "", fmt.Errorf("Error: expected 1 path, got 0. Usage: hexlet-path-size [options] <path>")
	}

	path := filepath.Clean(args[0])
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return pathsize.CLIArgs{}, "", fmt.Errorf("Error: Recived path \"%s\" is not exist", path)
		}
		return pathsize.CLIArgs{}, "", fmt.Errorf("Error with recived path \"%s\"", path)
	}

	return pathsize.Create(
		cmd.Bool("recursive"), 
		cmd.Bool("all"), 
		cmd.Bool("human"),
		), path, nil
}

func ActionLogic(ctx context.Context, cmd *cli.Command) error {
	cliArgs, path, err := cliArgsParse(cmd)
	if err != nil {
		return err
	}

	result, err := pathsize.Analyze(cliArgs, path)
	if err != nil {
		return fmt.Errorf("analyzing path %s: %w", path, err)
	}

	var formattedResult string
	if cliArgs.Human {
		formattedResult = pathsize.Humanity(result)
	} else {
		formattedResult = fmt.Sprintf("%dB", result)
	}
	

	fmt.Printf("%s\t%s\n", formattedResult, path)

	return nil
}

// Run CLI Util.

func Run() {
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
		Action: ActionLogic,
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

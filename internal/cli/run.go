// Package cli provides functions for run cli util.
package cli

import (
    "fmt"
    "log"
    "os"
    "context"
	"errors"

	"code/internal/analyzer"
	"code/internal/flags"
	"code/internal/formatter"
    "github.com/urfave/cli/v3"
)

// Error: path required
var ErrPathRequired = errors.New("path is required")


//nolint:forbidigo
// ActionLogic processes the CLI command and prints the file/directory size.
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

	if cliflags.Human {
		fmt.Printf("Directory %s has size: %s\n",path, formatter.Humanity(result))

		return nil
	}

	fmt.Println(result)

	return nil
}

// Run CLI Util.
//nolint:exhaustruct
func Run() {
    cmd := &cli.Command{
        Name:		"hexlet-path-size",
		Usage:		"print size of a file or directory",
		UsageText: 	"hexlet-path-size [global options] <path>",
		Flags: []cli.Flag{
        	&cli.BoolFlag{Name: "recursive", Aliases: []string{"r"}},
			&cli.BoolFlag{Name: "all", Aliases: []string{"a"}},
			&cli.BoolFlag{Name: "human"},
        },
		Action: ActionLogic,
    }

	err := cmd.Run(context.Background(), os.Args)
    if err != nil {
        log.Fatal(err)
    }
}
package cli

import (
    "fmt"
    "log"
    "os"
    "context"

	"code/internal/analyzer"
	"code/internal/flags"
	"code/internal/formatter"
    "github.com/urfave/cli/v3"
)

//Action Logic
func ActionLogic(ctx context.Context, cmd *cli.Command) error {
	if cmd.NArg() == 0 {
		return fmt.Errorf("path is required")
	}

	path := cmd.Args().First()
	cliflags := flags.Create(
		cmd.Bool("recursive"),
		cmd.Bool("all"),
		cmd.Bool("human"),
	)
	result, err := analyzer.Analyze(cliflags, path)
	if err != nil {
		return err
	}

	if cliflags.Human {
		fmt.Printf("Directory %s have size: %s\n",path, formatter.Humanity(result))
		return err
	}
	
	fmt.Println(result)
	return err
	
	
}

//Run CLI Util
func Run() {
    cmd := &cli.Command{
        Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		UsageText: "hexlet-path-size [global options] <path>",
		Flags: []cli.Flag{
        	&cli.BoolFlag{Name: "recursive", Aliases: []string{"r"}},
			&cli.BoolFlag{Name: "all", Aliases: []string{"a"}},
			&cli.BoolFlag{Name: "human"},
        },
		Action: ActionLogic,
    }

    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}
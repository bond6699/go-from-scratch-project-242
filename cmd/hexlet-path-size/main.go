package main

import (
    "context"
    "fmt"
    "log"
    "os"
    
    "github.com/urfave/cli/v3"
    "code/internal/pathsize"  // Импортируем пакет pathsize из модуля code
)

func Action(ctx context.Context, cmd *cli.Command) error {
    config := pathsize.Config{
        Path:      cmd.Args().First(),
        Recursive: cmd.Bool("recursive"),
        Human:     cmd.Bool("humanity"),
        All:       cmd.Bool("all"),
    }
    
    if config.Path == "" {
        return fmt.Errorf("path is required")
    }
    
    result, err := pathsize.GetPathSize(config)
    if err != nil {
        return fmt.Errorf("failed to get path size: %w", err)
    }
    
    fmt.Printf("%s\t%d bytes\n", config.Path, result)
    
    return nil
}

func main() {
    cmd := &cli.Command{
        Name:      "hexlet-path-size",
        Usage:     "print size of a file or directory",
        UsageText: "hexlet-path-size [options] <path>",
        Flags: []cli.Flag{
            &cli.BoolFlag{
                Name:    "humanity", 
				Aliases: []string{"hum"},
                Usage:   "print size in human-readable format (KB, MB, GB)",
            },
            &cli.BoolFlag{
                Name:    "recursive",
                Aliases: []string{"r"},
                Usage:   "process directories recursively",
            },
            &cli.BoolFlag{
                Name:    "all",
                Aliases: []string{"a"},
                Usage:   "include hidden files",
            },
        },
        Action: Action,
    }

    if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}
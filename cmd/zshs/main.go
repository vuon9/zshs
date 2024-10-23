package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/urfave/cli/v2"
	"github.com/vuon9/zshs/internal/zshs"
)

func main() {

	app := &cli.App{
		Name:  "zshs",
		Usage: "Search help for an alias in zsh plugins. For example: 'zshs git commit'",
		Commands: []*cli.Command{
			{
				Name:  "plugins",
				Usage: "Display all plugins as a checklist",
				Action: func(ctx *cli.Context) error {
					keyword := ctx.Args().Get(0)
					plugins, err := zshs.ListPlugins(os.Getenv("ZSH"))
					if err != nil {
						log.Fatal(err)
						cli.ShowAppHelpAndExit(ctx, 1)
					}

					if keyword != "" {
						plugins = zshs.FilterPlugins(plugins, keyword)
					}

					if len(plugins) > 0 {
						for _, plugin := range plugins {
							status := "❌"
							if plugin.Installed {
								status = "✅"
							}
							fmt.Printf("%s  %s\n", status, plugin.Name)
						}
					} else {
						fmt.Printf("Seems like '%s' is not one of official zsh plugins", keyword)
					}

					return nil
				},
			},
		},
		Action: func(ctx *cli.Context) error {
			if len(ctx.Args().Slice()) < 2 {
				cli.ShowAppHelpAndExit(ctx, 1)
			}

			plugin := ctx.Args().Get(0)
			keyword := ctx.Args().Get(1)
			result, err := zshs.SearchPluginCommandHelp(plugin, keyword, os.Getenv("ZSH"))
			if err != nil {
				log.Fatal(err)
				cli.ShowAppHelpAndExit(ctx, 1)
			}

			if len(result) > 0 {
				table := zshs.FormatAsMarkdownTable(result)
				out, _ := glamour.Render(table, "dark")
				fmt.Print(out)
			} else {
				fmt.Println("No results found.")
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

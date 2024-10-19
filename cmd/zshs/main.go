package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/urfave/cli/v2"
	"github.com/vuon9/zshs/package/zshs"
)

func main() {

	app := &cli.App{
		Name:  "zshs",
		Usage: "Search help for an alias in zsh plugins",
		Action: func(ctx *cli.Context) error {
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

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vuon9/zshs/package/zshs"
	"github.com/charmbracelet/glamour"
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
				panic(err)
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

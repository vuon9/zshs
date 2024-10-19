package zshs

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const pluginsFolder string = "/plugins"

type CommandHelp struct {
	Alias       string
	Command     string
	Description string
}

func SearchPluginCommandHelp(plugin string, keyword string, fromFolder string) ([]CommandHelp, error) {
	plugins := fromFolder + pluginsFolder
	dirs, err := os.ReadDir(plugins)
	if err != nil {
		panic(err)
	}

	var readmeFile string

	for _, dir := range dirs {
		if plugin == dir.Name() {
			readmeFile = fmt.Sprintf("%s/%s/%s/README.md", fromFolder, pluginsFolder, plugin)
			break
		}
	}

	if readmeFile == "" {
		return nil, fmt.Errorf("Can't find plugin %s", plugin)
	}

	readmeContent, err := os.ReadFile(readmeFile)
	if err != nil {
		return nil, err
	}

	matchedLines := []CommandHelp{}
	scanner := bufio.NewScanner(strings.NewReader(string(readmeContent)))

	for scanner.Scan() {
		line := scanner.Text()
		if matched, _ := regexp.MatchString(fmt.Sprintf("(?m)%s", keyword), line); matched {
			slc := strings.Split(line, "|")
			for j, slcp := range slc {
				slc[j] = strings.TrimSpace(slcp)
			}

			if len(slc) > 3 {
				matchedLines = append(matchedLines, CommandHelp{
					Alias:       slc[1],
					Command:     slc[2],
					Description: slc[3],
				})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return matchedLines, nil
}

func FormatAsMarkdownTable(results []CommandHelp) string {
	table := "| Zsh Alias | Real Command | Description |\n"
	table += "|-----------|--------------|-------------|\n"
	for _, item := range results {
		table += fmt.Sprintf("| %s | %s | %s |\n", item.Alias, item.Command, item.Description)
	}
	return table
}

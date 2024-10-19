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

type Plugin struct {
	Name      string
	Installed bool
}

func SearchPluginCommandHelp(plugin string, keyword string, fromFolder string) ([]*CommandHelp, error) {
	plugins, err := ListPlugins(fromFolder)
	if err != nil {
		return nil, err
	}

	var readmeFile string
	for _, pl := range plugins {
		if pl.Name == plugin {
			readmeFile = fmt.Sprintf("%s/%s/%s/README.md", fromFolder, pluginsFolder, pl.Name)
		}
	}

	if readmeFile == "" {
		return nil, fmt.Errorf("Couldn't find the plugin")
	}

	readmeContent, err := os.ReadFile(readmeFile)
	if err != nil {
		return nil, err
	}

	matchedLines := []*CommandHelp{}
	scanner := bufio.NewScanner(strings.NewReader(string(readmeContent)))

	for scanner.Scan() {
		line := scanner.Text()
		if matched, _ := regexp.MatchString(fmt.Sprintf("(?m)%s", keyword), line); matched {
			slc := strings.Split(line, "|")
			for j, slcp := range slc {
				slc[j] = strings.TrimSpace(slcp)
			}

			if len(slc) > 3 {
				matchedLines = append(matchedLines, &CommandHelp{
					Alias:       slc[1],
					Command:     slc[2],
					Description: slc[3],
				})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Couldn't read plugin README file")
	}

	return matchedLines, nil
}

func FormatAsMarkdownTable(results []*CommandHelp) string {
	table := "| Zsh Alias | Real Command | Description |\n"
	table += "|-----------|--------------|-------------|\n"
	for _, item := range results {
		table += fmt.Sprintf("| %s | %s | %s |\n", item.Alias, item.Command, item.Description)
	}
	return table
}

func ListPlugins(fromFolder string) ([]*Plugin, error) {
	plugins := fromFolder + pluginsFolder
	dirs, err := os.ReadDir(plugins)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	h := os.Getenv("HOME")
	rcFile, err := os.Open(fmt.Sprintf("%s/.zshrc", h))
	if err != nil {
		return nil, err
	}
	defer rcFile.Close()

	installedPluginsList := make([]string, 0)
	scanner := bufio.NewScanner(rcFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "plugins=(") && strings.HasSuffix(line, ")") {
			line = strings.Replace(line, "plugins=(", "", 1)
			line = strings.Replace(line, ")", "", 1)
			installedPluginsList = strings.Split(line, " ")
		}
	}

	mInstalledPlugins := make(map[string]bool)
	for _, installedPlugin := range installedPluginsList {
		mInstalledPlugins[installedPlugin] = true
	}

	pluginList := make([]*Plugin, len(dirs), len(dirs))
	for i, dir := range dirs {
		pluginList[i] = &Plugin{
			Name:      dir.Name(),
			Installed: mInstalledPlugins[dir.Name()],
		}
	}

	return pluginList, nil
}

func FilterPlugins(plugins []*Plugin, keyword string) []*Plugin {
	var filteredPlugins []*Plugin
	for _, plugin := range plugins {
		if strings.Contains(plugin.Name, keyword) {
			filteredPlugins = append(filteredPlugins, plugin)
		}
	}
	return filteredPlugins
}

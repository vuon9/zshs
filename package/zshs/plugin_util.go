package zshs

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const pluginsFolder string = "/plugins"

func SearchPluginCommandHelp(plugin string, keyword string, fromFolder string) ([]string, error) {
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

	matchedLines := []string{}
	scanner := bufio.NewScanner(strings.NewReader(string(readmeContent)))

	for scanner.Scan() {
		line := scanner.Text()
		if matched, _ := regexp.MatchString(fmt.Sprintf("(?m)%s", keyword), line); matched {
			matchedLines = append(matchedLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	for i, match := range matchedLines {
		slc := strings.Split(match, "|")
		for j, slcp := range slc {
			slc[j] = strings.TrimSpace(slcp)
		}

		if len(slc) > 3 {
			matchedLines[i] = fmt.Sprintf("%s - %s: %s", slc[1], slc[2], slc[3])
		} else {
			continue
		}
	}

	return matchedLines, nil
}

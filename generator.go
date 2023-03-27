package main

import (
	"fmt"
	"os"
	"strings"
)

var badges = []string{
	"license",
	"downloads",
	"last-commit",
	"repo-size",
	"stars",
}

func GenerateHeader(config Config) string {
	header := fmt.Sprintf("<div align=\"center\"><h1>%s</h1><p>%s<p>", config.Name, config.Description)
	if config.GhRepoPath != "" {
		header += "\n<div align=\"center\">"
		for _, badge := range badges {
			title := strings.Title(strings.ReplaceAll(badge, "-", ""))
			header += fmt.Sprintf("\n<img alt=%s src=\"https://img.shields.io/github/%s/%s?color=c3e7ff&style=flat-square\"", title, badge, config.GhRepoPath)
		}
		header += "\n</div>"
	}
	header += "</div>\n"
	return header
}

func GenerateList(title string, content []string) string {
	if len(content) == 0 {
		return ""
	}
	text := fmt.Sprintf("\n## %s\n", title)
	for _, listItem := range content {
		text += fmt.Sprintf("- [X] %s\n", listItem)
	}
	return text
}

func GenerateContributionGuidelines(config Config) string {
	repoUrl := fmt.Sprintf("https://github.com/%s/", config.GhRepoPath)
	return fmt.Sprintf(`%s## Feedback and contributions
***All contributions are very welcome!***

* Feel free to join the [dedicated community space](%s) for discussions about the app.
* Bug reports and feature requests can be submitted [here](%s) (please make sure to fill out all the requested information properly!).
* If you are a developer and wish to contribute to the app, please **fork** the project and submit a [**pull request**](%s).%s
`, "\n", config.CommunityUrl, repoUrl+"issues", repoUrl+"pulls", "\n")
}

func GenerateLicenseInfo(config Config) string {
	info := "\n## License\n\n"
	info += fmt.Sprintf("%s is licensed under the [%s](%s)", config.Name, config.LicenseName, config.LicenseUrl)
	return info
}

func Generate(config Config) string {
	content := GenerateHeader(config)
	content += GenerateList("Features", config.Features)
	content += GenerateList("Dependencies", config.Dependencies)
	content += GenerateContributionGuidelines(config)
	content += GenerateLicenseInfo(config)
	return content
}

func WriteFile(config Config) {
	content := Generate(config)
	os.WriteFile("README.md", []byte(content), os.ModePerm)
}

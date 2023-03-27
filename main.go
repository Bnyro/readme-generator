// Demo code for the Form primitive.
// Demo code for the Form primitive.
package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	config := Config{}

	app := tview.NewApplication()
	form := tview.NewForm().
		AddInputField("Project name", "", 20, nil, func(text string) {
			config.Name = text
		}).
		AddInputField("Description", "", 20, nil, func(text string) {
			config.Description = text
		}).
		AddInputField("Homepage", "", 20, nil, func(text string) {
			config.Homepage = text
		}).
		AddTextArea("Features", "", 40, 0, 0, func(text string) {
			config.Features = strings.Split(text, ",")
		}).
		AddTextArea("Dependencies", "", 40, 0, 0, func(text string) {
			config.Dependencies = strings.Split(text, ",")
		}).
		AddInputField("Author name", "", 20, nil, func(text string) {
			config.Author = text
		}).
		AddInputField("Author url", "", 20, nil, func(text string) {
			config.AuthorUrl = text
		}).
		AddInputField("License name", "", 20, nil, func(text string) {
			config.LicenseName = text
		}).
		AddInputField("License url", "", 20, nil, func(text string) {
			config.LicenseUrl = text
		}).
		AddInputField("Community url", "", 20, nil, func(text string) {
			config.CommunityUrl = text
		}).
		AddInputField("Gh repo path", "", 20, nil, func(text string) {
			config.GhRepoPath = text
		}).
		AddButton("Save", func() {
			WriteFile(config)
			app.Stop()
		}).
		AddButton("Quit", func() {
			app.Stop()
		})

	form.SetBorder(true).SetTitle("README generator").SetTitleAlign(tview.AlignLeft)
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			app.Stop()
		}
		return event
	})

	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

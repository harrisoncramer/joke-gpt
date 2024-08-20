package cmd

import (
	"fmt"

	help "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	keys     keyMap
	help     help.Model
	selector Selector
	err      error
}

func newFirstModel() tea.Model {
	return MainModel{
		keys:     newKeys(false),
		help:     help.New(),
		selector: newSelector(),
	}
}

func (m MainModel) Init() tea.Cmd {
	return m.getOptions
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case errMsg:
		m.err = msg

	/* Logic for the selector */
	case optionsMsg:
		m.selector.options = msg.options
	case moveMsg:
		m.selector.move(msg)
	case selectMsg:
		if msg.value == "Second" {
			secondModel := newSecondModel()
			return secondModel, secondModel.Init()
		}
	case tea.KeyMsg:
		switch msg.String() {
		case pluginOpts.Keys.Quit:
			return m, tea.Quit
		}
		return m, m.selector.Input(msg)
	}
	return m, nil
}

func (m MainModel) View() string {
	base := "Main View\n"
	base += m.selector.Render()
	base += fmt.Sprintf("\n\n%s", m.help.View(m.keys))
	return base
}

func (m MainModel) quit(msg tea.Msg) tea.Cmd {
	return quit(msg, m.keys.Quit)
}

func (m MainModel) getOptions() tea.Msg {
	return optionsMsg{
		options: []Option{
			{
				Label: "Blah",
				Value: "Blah",
			},
			{
				Label: "Second",
				Value: "Second",
			},
		},
	}
}

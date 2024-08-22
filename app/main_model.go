package app

import (
	"fmt"

	help "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	keys     keyMap
	help     help.Model
	selector tea.Model
	err      error
}

func NewFirstModel() tea.Model {
	return MainModel{
		keys: newKeys(false),
		help: help.New(),
		selector: Selector{
			options: []Option{
				{
					Label: "View Config",
					Value: "view_config",
				},
				{
					Label: "Second",
					Value: "second",
				},
			},
		},
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.err != nil {
		return m, tea.Quit
	}

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
	case optionsMsg, moveMsg:
		m.selector, cmd = m.selector.Update(msg)
	case selectMsg:
		if msg.value == "second" {
			secondModel := newSecondModel()
			return secondModel, secondModel.Init()
		}
	case tea.KeyMsg:
		switch msg.String() {
		case PluginOptions.Keys.Quit:
			return m, tea.Quit
		case PluginOptions.Keys.Down, PluginOptions.Keys.Up, PluginOptions.Keys.Select:
			m.selector, cmd = m.selector.Update(msg)
		}
	}
	return m, cmd
}

func (m MainModel) View() string {
	base := "Main View\n"
	base += m.selector.View()
	base += fmt.Sprintf("\n\n%s", m.help.View(m.keys))
	return base
}

func (m MainModel) quit(msg tea.Msg) tea.Cmd {
	return quit(msg, m.keys.Quit)
}

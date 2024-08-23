package app

import (
	"fmt"

	help "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	help     help.Model
	selector Selector
	err      error
}

var jokeOption = Option{
	Label: "Tell Joke",
	Value: "joke",
}

var quitOption = Option{
	Label: "Quit",
	Value: "quit",
}

func NewFirstModel() tea.Model {
	return MainModel{
		help: help.New(),
		selector: SelectorModel{
			options: []Option{jokeOption, quitOption},
		},
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	debugMsg(m, msg)
	if m.err != nil {
		return m, tea.Quit
	}

	/* Handle possible commands by selector */
	updatedSelector, selectorCmd := m.selector.Update(msg)
	m.selector = updatedSelector
	if selectorCmd != nil {
		return m, selectorCmd
	}

	/* All other events */
	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
		return m, nil
	case selectMsg:
		if msg.option.Value == jokeOption.Value {
			jokeModel := NewJokeModel()
			return jokeModel, jokeModel.Init()
		}
		if msg.option.Value == quitOption.Value {
			return m, tea.Quit
		}
	case tea.KeyMsg:
		switch msg.String() {
		case PluginOptions.Keys.Quit:
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m MainModel) View() string {
	base := "GPT Joke\n\n"
	base += m.selector.View()
	base += fmt.Sprintf("\n\n%s", m.help.View(newKeys(false)))
	return base
}

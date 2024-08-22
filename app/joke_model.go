package app

import (
	"fmt"

	help "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/harrisoncramer/joke-gpt/shared"
)

type JokeModel struct {
	err  error
	help help.Model
	keys keyMap
	joke string
}

type tellJokeMsg struct{}

func (m JokeModel) Init() tea.Cmd {
	return getJoke
}

func (m JokeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.err != nil {
		return m, tea.Quit
	}

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
	case jokeMsg:
		m.joke = msg.joke
	case tellJokeMsg:
		return m, getJoke
	case tea.KeyMsg:
		switch msg.String() {
		case PluginOptions.Keys.Repeat:
			m.joke = ""
			return m, getJoke
		case PluginOptions.Keys.Quit:
			return m, tea.Quit
		case PluginOptions.Keys.Back:
			firstModel := NewFirstModel(shared.AppStartArgs{})
			return firstModel, firstModel.Init()
		}
	}

	return m, cmd
}

func (m JokeModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	base := "GPT Joke - Joke View\n"

	if m.joke == "" {
		base += "Loading...\n"
	} else {
		base += fmt.Sprintf("\n%s", m.joke)
	}

	base += fmt.Sprintf("\n\n%s", m.help.View(newKeys(true)))

	return base
}

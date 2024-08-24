package app

import (
	"fmt"

	help "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/harrisoncramer/joke-gpt/shared"
)

type MainModel struct {
	help     help.Model
	selector Selector
	err      error
}

var jokeOption = Option{
	Label: "Tell Joke",
	Value: shared.JokeView,
}

var quitOption = Option{
	Label: "Quit",
	Value: "quit",
}

func NewMainModel() tea.Model {
	s := NewSelectorModel(NewSelectorModelOpts{
		filter: FilterOpts{
			placeholder: "Search...",
		},
		options: []Option{jokeOption, quitOption},
	})

	h := help.New()
	m := MainModel{
		help:     h,
		selector: s,
	}

	return Router{
		Model: m,
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

	var cmds = []tea.Cmd{}

	/* Handle possible commands by selector */
	var cmd tea.Cmd
	m.selector, cmd = m.selector.Update(msg)
	cmds = append(cmds, cmd)

	/* All other events */
	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
	case selectMsg:
		return m, changeView(msg.option.Value)
	case tea.KeyMsg:
		switch msg.String() {
		case PluginOptions.Keys.Quit:
			return m, tea.Quit
		case PluginOptions.Keys.Help:
			m.help.ShowAll = !m.help.ShowAll
		}
	}

	return m, tea.Batch(cmds...)
}

func (m MainModel) View() string {
	base := "GPT Joke\n\n"
	base += m.selector.View()
	base += fmt.Sprintf("\n\n%s", m.help.View(newKeys()))
	return base
}

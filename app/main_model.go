package app

import (
	"fmt"

	help "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/harrisoncramer/joke-gpt/app/router"
	"github.com/harrisoncramer/joke-gpt/shared"
)

type MainModel struct {
	help     help.Model
	selector Selector
	err      error
}

func NewMainModel() tea.Model {
	s := NewSelectorModel(NewSelectorModelOpts{
		filter: FilterOpts{
			placeholder: "Search...",
		},
		options: []Option{
			{"Tell Joke", shared.JokeView},
			{"Tell A different joke", shared.JokeViewTwo},
			{"Quit", "quit"},
		},
	})

	h := help.New()
	return MainModel{
		help:     h,
		selector: s,
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
		if msg.option.Value == "quit" {
			return m, tea.Quit
		}
		return m, router.ChangeView(msg.option.Value)
	case tea.KeyMsg:
		switch msg.String() {
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

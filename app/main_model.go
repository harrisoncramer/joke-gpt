package app

import (
	"fmt"

	help "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/harrisoncramer/joke-gpt/pkg/components"
	"github.com/harrisoncramer/joke-gpt/pkg/logger"
	"github.com/harrisoncramer/joke-gpt/pkg/router"
	"github.com/harrisoncramer/joke-gpt/shared"
)

type MainModel struct {
	help     help.Model
	selector components.Selector
	err      error
}

func NewMainModel() tea.Model {
	s := components.NewSelectorModel(components.NewSelectorModelOpts{
		Filter: components.FilterOpts{
			Placeholder: "Search...",
		},
		Options: []components.SelectorOption{
			{Label: "Tell Joke", Value: shared.JokeView},
			{Label: "Quit", Value: "quit"},
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
	logger.DebugMsg(m, msg)
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
	case components.SelectMsg:
		if msg.Option.Value == "quit" {
			return m, tea.Quit
		}
		return m, router.ChangeView(msg.Option.Value)
	case tea.KeyMsg:
		switch msg.String() {
		case shared.PluginOptions.Keys.Help:
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

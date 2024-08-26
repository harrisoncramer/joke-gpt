package app

import (
	"fmt"
	"os"

	help "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/harrisoncramer/joke-gpt/internal/logger"
	"github.com/harrisoncramer/joke-gpt/pkg/components"
	"github.com/harrisoncramer/joke-gpt/shared"
)

type MultiChoiceModel struct {
	help          help.Model
	multiSelector tea.Model
	err           error
}

func NewMultiChoiceModel() tea.Model {
	s := components.NewMultiSelectorModel(components.NewMultiSelectorModelOpts{
		Filter: components.FilterOpts{
			Placeholder: "Search...",
		},
		Options: []components.MultiSelectorOption{
			{Label: "Joke about bananas", Value: "bananas"},
			{Label: "Joke about chimpanzees", Value: "chimpanzees"},
			{Label: "Quit", Value: "quit"},
		},
	})

	h := help.New()
	return MultiChoiceModel{
		help:          h,
		multiSelector: s,
	}
}

func (m MultiChoiceModel) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, m.multiSelector.Init())
	return tea.Batch(cmds...)
}

func (m MultiChoiceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	logger.DebugMsg(m, msg)
	if m.err != nil {
		return m, tea.Quit
	}

	var cmds = []tea.Cmd{}

	/* Handle possible commands by selector */
	var cmd tea.Cmd
	m.multiSelector, cmd = m.multiSelector.Update(msg)
	cmds = append(cmds, cmd)

	/* All other events */
	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
	case components.MultiSelectMsg:
		fmt.Printf("Not implemented")
		os.Exit(0)
	case tea.KeyMsg:
		switch msg.String() {
		case shared.PluginOptions.Keys.Help:
			m.help.ShowAll = !m.help.ShowAll
		}
	}

	return m, tea.Batch(cmds...)
}

func (m MultiChoiceModel) View() string {
	base := "GPT Multi Choice\n\n"
	base += m.multiSelector.View()
	base += fmt.Sprintf("\n\n%s", m.help.View(newKeys()))
	return base
}

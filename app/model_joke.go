package app

import (
	"fmt"

	help "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/harrisoncramer/joke-gpt/internal/logger"
	"github.com/harrisoncramer/joke-gpt/pkg/router"
	"github.com/harrisoncramer/joke-gpt/shared"
)

type JokeModel struct {
	err     error
	help    help.Model
	keys    keyMap
	joke    string
	spinner spinner.Model
}

func NewJokeModel() tea.Model {
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return JokeModel{
		help:    help.New(),
		keys:    newKeys(),
		spinner: s,
	}
}

type tellJokeMsg struct{}

func (m JokeModel) Init() tea.Cmd {
	var cmds = []tea.Cmd{}
	cmds = append(cmds, getJoke, m.spinner.Tick)
	return tea.Batch(cmds...)
}

func (m JokeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	logger.DebugMsg(m, msg)
	if m.err != nil {
		return m, tea.Quit
	}

	cmds := []tea.Cmd{}

	/* Handle messages from the spinner */
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
	case jokeMsg:
		m.joke = msg.joke
	case tellJokeMsg:
		cmds = append(cmds, getJoke)
	case tea.KeyMsg:
		switch msg.String() {
		case shared.PluginOptions.Keys.Repeat:
			m.joke = ""
			cmds = append(cmds, getJoke)
		case shared.PluginOptions.Keys.Back:
			return m, router.Back()
		}
	}

	return m, tea.Batch(cmds...)
}

func (m JokeModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	base := "GPT Joke - Joke View\n"

	if m.joke == "" {
		base += fmt.Sprintf("\n%s", m.spinner.View())
	} else {
		base += fmt.Sprintf("\n%s", m.joke)
	}

	base += fmt.Sprintf("\n\n%s", m.help.View(newKeys()))

	return base
}

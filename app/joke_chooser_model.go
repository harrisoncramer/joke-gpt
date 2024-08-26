package app

import (
	"fmt"

	help "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/harrisoncramer/joke-gpt/internal/logger"
	"github.com/harrisoncramer/joke-gpt/pkg/components"
	"github.com/harrisoncramer/joke-gpt/shared"
)

type MultiChoiceModel struct {
	help         help.Model
	choosingJoke bool
	jokeChooser  tea.Model
	voiceChooser tea.Model
	spinner      spinner.Model
	loading      bool
	result       string
	jokes        components.MultiSelectorOptions
	err          error
}

func NewMultiChoiceModel() tea.Model {
	return MultiChoiceModel{
		help:         help.New(),
		choosingJoke: true,
		jokeChooser: components.NewMultiSelectorModel(components.NewMultiSelectorModelOpts{
			Filter: components.FilterOpts{
				Placeholder: "Search...",
			},
			Options: components.MultiSelectorOptions{
				{Label: "Joke about bananas", Value: "bananas"},
				{Label: "Joke about chimpanzees", Value: "chimpanzees"},
			},
		}),
		voiceChooser: components.NewSelectorModel(components.NewSelectorModelOpts{
			Filter: components.FilterOpts{
				Placeholder: "Search...",
			},
			Options: components.SelectorOptions{
				{Label: "Old-timey", Value: "a 1950s man, with a transatlantic radio style"},
				{Label: "Futuristic", Value: "a robot from the future"},
			},
		}),
		spinner: spinner.New(),
	}
}

func (m MultiChoiceModel) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, m.jokeChooser.Init(), m.voiceChooser.Init())
	return tea.Batch(cmds...)
}

func (m MultiChoiceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	logger.DebugMsg(m, msg)
	if m.err != nil {
		return m, tea.Quit
	}

	var cmds = []tea.Cmd{}

	/* Handle possible commands by joke chooser */
	var cmd tea.Cmd
	m.jokeChooser, cmd = m.jokeChooser.Update(msg)
	cmds = append(cmds, cmd)

	/* Possible commands by ai chooser */
	m.voiceChooser, cmd = m.voiceChooser.Update(msg)
	cmds = append(cmds, cmd)

	/* Handle messages from the spinner */
	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)

	/* All other events */
	switch msg := msg.(type) {
	case errMsg:
		m.loading = false
		m.err = msg
	case jokeMsg:
		m.loading = false
		m.result = msg.joke
	case components.MultiSelectMsg:
		m.jokes = msg.Options
		m.choosingJoke = false
	case components.SelectMsg:
		if len(m.jokes) > 0 {
			m.loading = true

			user := "Tell me one joke about each of the following subjects: "
			for _, joke := range m.jokes {
				user += fmt.Sprintf("%s,", joke.Value)
			}

			system := fmt.Sprintf("You are a a wisecracking assistant with a voice in the style of %s.", msg.Option.Value)
			cmds = append(cmds, runPrompt(Prompt{
				user:   user,
				system: system,
			}), m.spinner.Tick)
		}
	case tea.KeyMsg:
		switch msg.String() {
		case shared.PluginOptions.Keys.Help:
			m.help.ShowAll = !m.help.ShowAll
		}
	}

	return m, tea.Batch(cmds...)
}

func (m MultiChoiceModel) View() string {
	base := ""

	if m.result == "" && m.loading {
		base += fmt.Sprintf("\n%s", m.spinner.View())
		return base
	} else if m.result != "" {
		base += fmt.Sprintf("\n%s", m.result)
		return base
	}

	if m.choosingJoke {
		base += "Select jokes\n\n"
		base += m.jokeChooser.View()
	} else {
		base += "Select Voice\n\n"
		base += m.voiceChooser.View()
	}
	base += fmt.Sprintf("\n\n%s", m.help.View(newKeys()))
	return base
}

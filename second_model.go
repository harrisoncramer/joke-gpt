package main

import (
	"errors"
	"fmt"

	help "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type SecondModel struct {
	err      error
	keys     keyMap
	help     help.Model
	selector Selector
}

func newSecondModel() NestedView {
	selector := newSelector(selectorOpts{url: "http://localhost:3000/options"})
	return SecondModel{
		keys: keyMap{
			Quit:   quitKeys,
			Back:   backKeys,
			Up:     selector.keys.Up,
			Down:   selector.keys.Down,
			Select: selector.keys.Select,
		},
		selector: selector,
		help:     help.New(),
	}
}

func (m SecondModel) Init() tea.Cmd {
	return m.selector.getOptions
}

func (m SecondModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	quit := m.quit(msg)
	if quit != nil {
		return m, quit
	}

	back := m.back(msg)
	if back != nil {
		return back, nil
	}

	switch msg := msg.(type) {
	case errMsg:
		m.err = msg

	/* Logic for the selector */
	case optionsMsg:
		m.selector.options = msg.options
	case move:
		m.selector.move(msg)
	case selectedEntry:
		m.err = errors.New("Chosen!")
		return m, nil
	case tea.KeyMsg:
		return m, m.selector.Input(msg)
	}

	return m, nil
}

func (m SecondModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	base := ""
	base += m.selector.Render()
	base += fmt.Sprintf("\n%s", m.help.View(m.keys))

	return base
}

func (m SecondModel) quit(msg tea.Msg) tea.Cmd {
	return quit(msg, m.keys.Quit)
}

func (m SecondModel) back(msg tea.Msg) Quitter {
	return back(msg, m.keys.Back)
}

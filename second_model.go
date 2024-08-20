package main

import (
	"fmt"
	"slices"

	help "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type SecondModel struct {
	keys keyMap
	help help.Model
}

func newSecondModel() NestedView {
	m := SecondModel{
		keys: keyMap{
			Quit: quitKey,
			Back: backKey,
		},
		help: help.New(),
	}
	return m
}

func (m SecondModel) Init() tea.Cmd {
	return nil
}

func (m SecondModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	quit := m.quit(m)
	if quit != nil {
		return m, quit
	}

	back := m.back(msg)
	if back != nil {
		return back, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if slices.Contains(m.keys.Quit.Keys(), msg.String()) {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m SecondModel) View() string {
	base := "Second view"
	helpView := m.help.View(m.keys)
	base += fmt.Sprintf("\n\n%s", helpView)
	return base
}

/* Non-Standard Methods */
func (m SecondModel) quit(msg tea.Msg) tea.Cmd {
	return quit(msg, m.keys)
}

func (m SecondModel) back(msg tea.Msg) Quitter {
	return back(msg, m.keys)
}

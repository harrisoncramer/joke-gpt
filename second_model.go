package main

import (
	"fmt"
	"slices"

	help "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Quit key.Binding
	Back key.Binding
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit},
	}
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Quit}
}

func newSecondModel() NestedView {
	m := SecondModel{
		keys: keyMap{
			Quit: key.NewBinding(
				key.WithKeys("ctrl+c"),
				key.WithHelp("q", "quit"),
			),
			Back: key.NewBinding(
				key.WithKeys("esc"),
				key.WithHelp("esc", "back"),
			),
		},
		help: help.New(),
	}
	return m
}

type SecondModel struct {
	keys keyMap
	help help.Model
}

func (m SecondModel) Init() tea.Cmd {
	return nil
}

func (m SecondModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	quit := m.quit(m)
	if quit != nil {
		return nil, quit
	}

	back := m.back(msg)
	if back != nil {
		return back, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return newFirstModel(), nil
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
	return quit(msg)
}

func (m SecondModel) back(msg tea.Msg) Quitter {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if slices.Contains(m.keys.Quit.Keys(), msg.String()) {
			return newFirstModel()
		}
	}
	return nil
}

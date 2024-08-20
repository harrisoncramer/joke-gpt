package main

import (
	"slices"

	tea "github.com/charmbracelet/bubbletea"
)

func newSecondModel() NestedView {
	m := SecondModel{
		backKeys: []string{"esc"},
	}
	return m
}

type SecondModel struct {
	backKeys []string
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
	return "Second view!"
}

/* Non-Standard Methods */
func (m SecondModel) quit(msg tea.Msg) tea.Cmd {
	return quit(msg)
}

func (m SecondModel) back(msg tea.Msg) Quitter {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if slices.Contains(m.backKeys, msg.String()) {
			return newFirstModel()
		}
	}
	return nil
}

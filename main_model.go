package main

import tea "github.com/charmbracelet/bubbletea"

func newFirstModel() Quitter {
	m := MainModel{}
	return m
}

type MainModel struct{}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := m.quit(msg)
	if cmd != nil {
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return newSecondModel().Update(msg)
		}
	}
	return m, nil
}

func (m MainModel) View() string {
	return "Hello there"
}

func (m MainModel) quit(msg tea.Msg) tea.Cmd {
	return quit(msg)
}

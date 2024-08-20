package main

import tea "github.com/charmbracelet/bubbletea"

type MainModel struct{}

func (m MainModel) handleExitKeys(msg tea.KeyMsg) tea.Cmd {
	switch msg.Type {
	case tea.KeyCtrlC, tea.KeyCtrlD:
		return tea.Quit
	}
	return nil
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:

		cmd := m.handleExitKeys(msg)
		if cmd != nil {
			return m, cmd
		}

		switch msg.String() {
		case "up", "k":
			return m, tea.Quit
		case "down", "j":
			return m, tea.Quit
		case "enter":
		}
	}
	return m, nil
}

func (m MainModel) View() string {
	return "Hello there"
}

package main

import tea "github.com/charmbracelet/bubbletea"

type QuittableModel interface {
	tea.Model
	handleExitKeys(msg tea.KeyMsg) tea.Cmd
}

package main

import tea "github.com/charmbracelet/bubbletea"

/* Handles quitting the application when certain keys are pressed */
func quit(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyCtrlD:
			return tea.Quit
		}
	}
	return nil
}

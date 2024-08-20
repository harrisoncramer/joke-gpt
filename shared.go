package main

import (
	"slices"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Quit key.Binding
	Back key.Binding
}

var quitKey = key.NewBinding(
	key.WithKeys(
		tea.KeyCtrlC.String(),
		tea.KeyCtrlD.String(),
	),
	key.WithHelp("ctrl+c/ctrl+d", "quit"),
)

var backKey = key.NewBinding(
	key.WithKeys(tea.KeyEsc.String()),
	key.WithHelp("esc", "back"),
)

/* Handles quitting the application when certain keys are pressed */
func quit(msg tea.Msg, keys keyMap) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if slices.Contains(keys.Quit.Keys(), msg.String()) {
			return tea.Quit
		}
	}
	return nil
}

/* Navigates to the previous model */
func back(msg tea.Msg, keys keyMap) Quitter {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if slices.Contains(keys.Back.Keys(), msg.String()) {
			return newFirstModel()
		}
	}
	return nil
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

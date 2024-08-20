package main

import (
	"slices"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Quit   key.Binding
	Back   key.Binding
	Select key.Binding
	Up     key.Binding
	Down   key.Binding
}

var quitKeys = key.NewBinding(
	key.WithKeys(tea.KeyCtrlC.String()),
	key.WithHelp("ctrl+c", "quit"),
)

var selectKeys = key.NewBinding(
	key.WithKeys(tea.KeyEnter.String()),
	key.WithHelp("enter", "select"),
)

var backKeys = key.NewBinding(
	key.WithKeys(tea.KeyEsc.String()),
	key.WithHelp("esc", "back"),
)

var upKeys = key.NewBinding(
	key.WithKeys("k"),
	key.WithHelp("k", "up"),
)

var downKeys = key.NewBinding(
	key.WithKeys("j"),
	key.WithHelp("j", "down"),
)

/* Handles quitting the application when certain keys are pressed */
func quit(msg tea.Msg, keybinding key.Binding) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if slices.Contains(keybinding.Keys(), msg.String()) {
			return tea.Quit
		}
	}
	return nil
}

/* Navigates to the previous model */
func back(msg tea.Msg, keybinding key.Binding) Quitter {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if slices.Contains(keybinding.Keys(), msg.String()) {
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
	return []key.Binding{
		k.Back,
		k.Quit,
		k.Select,
		k.Up,
		k.Down,
	}
}

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

package cmd

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

func newKeys(isNested bool) keyMap {
	k := keyMap{
		Quit: key.NewBinding(
			key.WithKeys(pluginOpts.Keys.Quit),
			key.WithHelp(pluginOpts.Keys.Quit, "quit"),
		),
		Select: key.NewBinding(
			key.WithKeys(pluginOpts.Keys.Select),
			key.WithHelp(pluginOpts.Keys.Select, "select"),
		),
		Up: key.NewBinding(
			key.WithKeys(pluginOpts.Keys.Up),
			key.WithHelp(pluginOpts.Keys.Up, "up"),
		),
		Down: key.NewBinding(
			key.WithKeys(pluginOpts.Keys.Down),
			key.WithHelp(pluginOpts.Keys.Down, "down"),
		),
	}

	if isNested {
		k.Back = key.NewBinding(
			key.WithKeys(pluginOpts.Keys.Back),
			key.WithHelp(pluginOpts.Keys.Back, "back"),
		)
	}

	return k
}

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

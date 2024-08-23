package app

import (
	"github.com/charmbracelet/bubbles/key"
)

type keyMap struct {
	Quit   key.Binding
	Back   key.Binding
	Select key.Binding
	Up     key.Binding
	Down   key.Binding
	Repeat key.Binding
	Filter key.Binding
	Help   key.Binding
}

func newKeys() keyMap {
	k := keyMap{
		Quit: key.NewBinding(
			key.WithKeys(PluginOptions.Keys.Quit),
			key.WithHelp(PluginOptions.Keys.Quit, "quit"),
		),
		Select: key.NewBinding(
			key.WithKeys(PluginOptions.Keys.Select),
			key.WithHelp(PluginOptions.Keys.Select, "select"),
		),
		Up: key.NewBinding(
			key.WithKeys(PluginOptions.Keys.Up),
			key.WithHelp(PluginOptions.Keys.Up, "up"),
		),
		Down: key.NewBinding(
			key.WithKeys(PluginOptions.Keys.Down),
			key.WithHelp(PluginOptions.Keys.Down, "down"),
		),
		Repeat: key.NewBinding(
			key.WithKeys(PluginOptions.Keys.Repeat),
			key.WithHelp(PluginOptions.Keys.Repeat, "repeat"),
		),
		Back: key.NewBinding(
			key.WithKeys(PluginOptions.Keys.Back),
			key.WithHelp(PluginOptions.Keys.Back, "back"),
		),
		Filter: key.NewBinding(
			key.WithKeys(PluginOptions.Keys.Filter),
			key.WithHelp(PluginOptions.Keys.Filter, "filter"),
		),
		Help: key.NewBinding(
			key.WithKeys(PluginOptions.Keys.Help),
			key.WithHelp(PluginOptions.Keys.Help, "toggle help"),
		),
	}

	return k
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.Back,
			k.Quit,
			k.Select,
			k.Up,
			k.Down,
			k.Repeat,
			k.Filter,
			k.Help,
		},
	}
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Quit,
		k.Help,
	}
}

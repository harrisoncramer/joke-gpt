package app

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/harrisoncramer/joke-gpt/shared"
)

type keyMap struct {
	Quit   key.Binding
	Back   key.Binding
	Select key.Binding
	Toggle key.Binding
	Up     key.Binding
	Down   key.Binding
	Repeat key.Binding
	Filter key.Binding
	Help   key.Binding
}

func newKeys() keyMap {
	k := keyMap{
		Quit: key.NewBinding(
			key.WithKeys(shared.PluginOptions.Keys.Quit),
			key.WithHelp(shared.PluginOptions.Keys.Quit, "quit"),
		),
		Select: key.NewBinding(
			key.WithKeys(shared.PluginOptions.Keys.Select),
			key.WithHelp(shared.PluginOptions.Keys.Select, "select"),
		),
		Toggle: key.NewBinding(
			key.WithKeys(shared.PluginOptions.Keys.Toggle),
			key.WithHelp(shared.PluginOptions.Keys.Toggle, "toggle"),
		),
		Up: key.NewBinding(
			key.WithKeys(shared.PluginOptions.Keys.Up),
			key.WithHelp(shared.PluginOptions.Keys.Up, "up"),
		),
		Down: key.NewBinding(
			key.WithKeys(shared.PluginOptions.Keys.Down),
			key.WithHelp(shared.PluginOptions.Keys.Down, "down"),
		),
		Repeat: key.NewBinding(
			key.WithKeys(shared.PluginOptions.Keys.Repeat),
			key.WithHelp(shared.PluginOptions.Keys.Repeat, "repeat"),
		),
		Back: key.NewBinding(
			key.WithKeys(shared.PluginOptions.Keys.Back),
			key.WithHelp(shared.PluginOptions.Keys.Back, "back"),
		),
		Filter: key.NewBinding(
			key.WithKeys(shared.PluginOptions.Keys.Filter),
			key.WithHelp(shared.PluginOptions.Keys.Filter, "filter"),
		),
		Help: key.NewBinding(
			key.WithKeys(shared.PluginOptions.Keys.Help),
			key.WithHelp(shared.PluginOptions.Keys.Help, "toggle help"),
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
			k.Toggle,
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

package cmd

import (
	"fmt"
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Selector struct {
	keys    keyMap
	cursor  int
	options []Option
}

type Direction string

const (
	Up   Direction = "up"
	Down Direction = "down"
)

type move struct {
	direction Direction
}

type selectedEntry struct {
	value string
}

func newSelector() Selector {

	keys := newKeys()

	m := Selector{
		cursor: 0,
		keys: keyMap{
			Select: keys.Select,
			Up:     keys.Up,
			Down:   keys.Down,
		},
	}

	return m
}

func (s *Selector) move(movement move) {
	if movement.direction == Up {
		if s.cursor > 0 {
			s.cursor--
		}
	} else {
		if s.cursor < len(s.options)-1 {
			s.cursor++
		}
	}
}

func (s Selector) Render() string {
	base := ""
	for i, option := range s.options {
		if i == s.cursor {
			base += fmt.Sprintf("%s %s\n", pluginOpts.Display.Cursor, option.Label)
		} else {
			base += fmt.Sprintf("%s %s\n", strings.Repeat(" ", len(pluginOpts.Display.Cursor)), option.Label)
		}
	}
	return base
}

func (s Selector) Input(msg tea.KeyMsg) tea.Cmd {
	return func() tea.Msg {
		str := msg.String()
		if slices.Contains(s.keys.Down.Keys(), str) {
			return move{direction: Down}
		}
		if slices.Contains(s.keys.Up.Keys(), str) {
			return move{direction: Up}
		}
		if slices.Contains(s.keys.Select.Keys(), str) {
			return selectedEntry{value: s.options[s.cursor].Value}
		}
		return nil
	}
}

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type OptionsResponse []Option

type optionsMsg struct {
	options []Option
}

package app

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Selector struct {
	cursor  int
	options []Option
}

type Direction string

const (
	Up   Direction = "up"
	Down Direction = "down"
)

type moveMsg struct {
	direction Direction
}

type selectMsg struct {
	value string
}

func newSelector() Selector {
	m := Selector{
		cursor: 0,
	}

	return m
}

func (s *Selector) move(movement moveMsg) {
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
			base += fmt.Sprintf("%s %s\n", PluginOptions.Display.Cursor, option.Label)
		} else {
			base += fmt.Sprintf("%s %s\n", strings.Repeat(" ", len(PluginOptions.Display.Cursor)), option.Label)
		}
	}
	return base
}

func (s Selector) Input(msg tea.KeyMsg) tea.Cmd {
	return func() tea.Msg {
		str := msg.String()
		if PluginOptions.Keys.Down == str {
			return moveMsg{direction: Down}
		}
		if PluginOptions.Keys.Up == str {
			return moveMsg{direction: Up}
		}
		if PluginOptions.Keys.Select == str {
			return selectMsg{value: s.options[s.cursor].Value}
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

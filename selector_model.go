package main

import (
	"fmt"
	"slices"

	tea "github.com/charmbracelet/bubbletea"
)

type Cursor struct {
	line int
}

type Selector struct {
	keys    keyMap
	cursor  Cursor
	options []Option
}

type up struct{}
type down struct{}
type selectedEntry struct{}

func newSelector() Selector {
	m := Selector{
		cursor: Cursor{},
		keys: keyMap{
			Select: selectKeys,
			Up:     upKeys,
			Down:   downKeys,
		},
	}
	return m
}

func selectEntry() tea.Msg {
	return selectedEntry{}
}

func (s *Selector) up() tea.Msg {
	if s.cursor.line > 0 {
		s.cursor.line--
	}
	return up{}
}

func (s *Selector) down() tea.Msg {
	if s.cursor.line < len(s.options)-1 {
		s.cursor.line++
	}
	return down{}
}

func (s Selector) Render() string {
	base := ""
	for i, option := range s.options {
		if i == s.cursor.line {
			base += fmt.Sprintf("> %s\n", option.Label)
		} else {
			base += fmt.Sprintf("  %s\n", option.Label)
		}
	}
	return base
}

func (s Selector) Input(msg tea.KeyMsg) tea.Cmd {
	str := msg.String()
	if slices.Contains(s.keys.Down.Keys(), str) {
		return s.down
	}
	if slices.Contains(s.keys.Up.Keys(), str) {
		return s.up
	}
	if slices.Contains(s.keys.Select.Keys(), str) {
		return selectEntry
	}

	return nil
}

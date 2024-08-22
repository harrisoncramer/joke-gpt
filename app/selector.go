package app

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

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

type optionsMsg struct {
	options []Option
}

func (s Selector) Init() tea.Cmd {
	return nil
}

/* Responds to keypresses that are defined in our plugin options and updates the model, and/or selects a value */
func (s Selector) Update(msg tea.Msg) (Selector, tea.Cmd) {
	switch msg := msg.(type) {
	case optionsMsg:
		s.setOptions(msg.options)
	case tea.KeyMsg:
		switch msg.String() {
		case PluginOptions.Keys.Down:
			s.move(Down)
		case PluginOptions.Keys.Up:
			s.move(Up)
		case PluginOptions.Keys.Select:
			return s, s.selectVal
		}
	}
	return s, nil
}

/* Renders the choices and the current cursor */
func (s Selector) View() string {
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

/* Moves the cursor up or down among the options */
func (s *Selector) move(direction Direction) {
	if direction == Up {
		if s.cursor > 0 {
			s.cursor--
		}
	} else {
		if s.cursor < len(s.options)-1 {
			s.cursor++
		}
	}
}

/* Sets options on the selector */
func (s *Selector) setOptions(options []Option) {
	s.options = options
}

/* Chooses the value at the given index */
func (s Selector) selectVal() tea.Msg {
	return selectMsg{value: s.options[s.cursor].Value}
}

/* If a key was pressed, call the update function if it's relevant. This lets us group all of the key logic in one method */
func (s Selector) maybeUpdate(msg tea.Msg) (Selector, tea.Cmd) {
	switch msg := msg.(type) {
	case optionsMsg, moveMsg:
		return s.Update(msg)
	case tea.KeyMsg:
		switch msg.String() {
		case PluginOptions.Keys.Down, PluginOptions.Keys.Up, PluginOptions.Keys.Select:
			return s.Update(msg)
		}
	}

	return s, nil
}

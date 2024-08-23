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

type Selector interface {
	tea.Model
}

type SelectorModel struct {
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
	option Option
}

type optionsMsg struct {
	options []Option
}

func (m SelectorModel) Init() tea.Cmd {
	return nil
}

/* Responds to keypresses that are defined in our plugin options and updates the model, and/or selects a value */
func (m SelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	debugMsg(m, msg)
	switch msg := msg.(type) {
	case optionsMsg:
		m.setOptions(msg.options)
	case tea.KeyMsg:
		switch msg.String() {
		case PluginOptions.Keys.Down:
			m.move(Down)
		case PluginOptions.Keys.Up:
			m.move(Up)
		case PluginOptions.Keys.Select:
			return m, m.selectVal
		}
	}
	return m, nil
}

/* Renders the choices and the current cursor */
func (m SelectorModel) View() string {
	base := ""
	for i, option := range m.options {
		if i == m.cursor {
			base += fmt.Sprintf("%s %s\n", PluginOptions.Display.Cursor, option.Label)
		} else {
			base += fmt.Sprintf("%s %s\n", strings.Repeat(" ", len(PluginOptions.Display.Cursor)), option.Label)
		}
	}
	return base
}

/* Moves the cursor up or down among the options */
func (m *SelectorModel) move(direction Direction) {
	if direction == Up {
		if m.cursor > 0 {
			m.cursor--
		}
	} else {
		if m.cursor < len(m.options)-1 {
			m.cursor++
		}
	}
}

/* Sets options on the selector */
func (m *SelectorModel) setOptions(options []Option) {
	m.options = options
}

/* Chooses the value at the given index */
func (s SelectorModel) selectVal() tea.Msg {
	return selectMsg{s.options[s.cursor]}
}

/* If a key was pressed, call the update function if it's relevant. This lets us group all of the key logic in one method */
func (s SelectorModel) maybeUpdate(msg tea.Msg) (Selector, tea.Cmd) {
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

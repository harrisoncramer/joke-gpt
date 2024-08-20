package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"time"

	help "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type Cursor struct {
	line int
}

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type OptionsResponse []Option

type optionsMsg struct {
	options []Option
}

type SecondModel struct {
	err     error
	keys    keyMap
	help    help.Model
	cursor  Cursor
	options []Option
}

func newSecondModel() NestedView {
	m := SecondModel{
		cursor: Cursor{},
		keys: keyMap{
			Quit:   quitKeys,
			Back:   backKeys,
			Select: selectKeys,
			Up:     upKeys,
			Down:   downKeys,
		},
		help: help.New(),
	}
	return m
}

var count = ""

func (m SecondModel) Init() tea.Cmd {
	count += "a"
	return checkServer
}

func (m SecondModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	quit := m.quit(msg)
	if quit != nil {
		return m, quit
	}

	back := m.back(msg)
	if back != nil {
		return back, nil
	}

	switch msg := msg.(type) {
	case up:
		if m.cursor.line > 0 {
			m.cursor.line--
		}
		return m, nil
	case down:
		if m.cursor.line < len(m.options)-1 {
			m.cursor.line++
		}
		return m, nil
	case statusMsg:
		return m, tea.Quit
	case errMsg:
		m.err = msg
		return m, nil
	case optionsMsg:
		m.options = msg.options
		return m, nil
	case tea.KeyMsg:
		str := msg.String()
		if slices.Contains(m.keys.Down.Keys(), str) {
			return m, moveDown
		}
		if slices.Contains(m.keys.Up.Keys(), str) {
			return m, moveUp
		}
	}

	return m, nil
}

func (m SecondModel) View() string {
	base := ""
	if m.err != nil {
		return m.err.Error()
	}

	for i, option := range m.options {
		if i == m.cursor.line {
			base += fmt.Sprintf("> %s\n", option.Label)
		} else {
			base += fmt.Sprintf("  %s\n", option.Label)
		}
	}

	base += fmt.Sprintf("\n%s", m.help.View(m.keys))

	return base
}

func (m SecondModel) quit(msg tea.Msg) tea.Cmd {
	return quit(msg, m.keys.Quit)
}

func (m SecondModel) back(msg tea.Msg) Quitter {
	return back(msg, m.keys.Back)
}

func checkServer() tea.Msg {
	c := &http.Client{Timeout: 10 * time.Second}
	res, err := c.Get("http://localhost:3000/options")

	if err != nil {
		return errMsg{err}
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return errMsg{err}
	}

	var optionsResponse []Option
	err = json.Unmarshal(data, &optionsResponse)
	if err != nil {
		return errMsg{err}
	}

	return optionsMsg{
		options: optionsResponse,
	}
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	help "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type OptionsResponse []Option

type optionsMsg struct {
	options []Option
}

type SecondModel struct {
	err      error
	keys     keyMap
	help     help.Model
	options  []Option
	selector Selector
}

func newSecondModel() NestedView {
	selector := newSelector()
	return SecondModel{
		keys: keyMap{
			Quit:   quitKeys,
			Back:   backKeys,
			Select: selector.keys.Select,
			Up:     selector.keys.Up,
			Down:   selector.keys.Down,
		},
		selector: selector,
		help:     help.New(),
	}
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
		m.selector.up()
	case down:
		m.selector.down()
	case statusMsg:
		return m, tea.Quit
	case errMsg:
		m.err = msg
	case optionsMsg:
		m.selector.options = msg.options
	case selectedEntry:
		m.err = errors.New("Chosen!")
		return m, nil
	case tea.KeyMsg:
		selectorChange := m.selector.Input(msg)
		if selectorChange != nil {
			return m, selectorChange
		}
	}

	return m, nil
}

func (m SecondModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	base := ""
	base += m.selector.Render()
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

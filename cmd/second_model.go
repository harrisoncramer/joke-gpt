package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"

	help "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type SecondModel struct {
	err      error
	keys     keyMap
	help     help.Model
	selector Selector
}

func newSecondModel() NestedView {
	selector := newSelector()
	return SecondModel{
		keys:     newKeys(),
		selector: selector,
		help:     help.New(),
	}
}

func (m SecondModel) Init() tea.Cmd {
	return m.getOptions
}

func (m SecondModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	quit := m.quit(msg)
	if quit != nil {
		return m, quit
	}

	switch msg := msg.(type) {
	case errMsg:
		m.err = msg

	/* Logic for the selector */
	case optionsMsg:
		m.selector.options = msg.options
	case move:
		m.selector.move(msg)
	case tea.KeyMsg:
		if slices.Contains(m.keys.Back.Keys(), msg.String()) {
			firstModel := newFirstModel()
			return firstModel, firstModel.Init()
		}
		return m, m.selector.Input(msg)
	}

	return m, nil
}

func (m SecondModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	base := "Second View\n"
	base += m.selector.Render()
	base += fmt.Sprintf("\n%s", m.help.View(m.keys))

	return base
}

func (m SecondModel) quit(msg tea.Msg) tea.Cmd {
	return quit(msg, m.keys.Quit)
}

func (m SecondModel) back(msg tea.Msg) Quitter {
	return newFirstModel()
}

func (m SecondModel) getOptions() tea.Msg {
	c := &http.Client{Timeout: pluginOpts.Network.TimeoutMillis}
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

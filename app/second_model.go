package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

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
	return SecondModel{
		keys:     newKeys(true),
		selector: newSelector(),
		help:     help.New(),
	}
}

func (m SecondModel) Init() tea.Cmd {
	return m.getOptions
}

func (m SecondModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.err != nil {
		return m, tea.Quit
	}

	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
	/* Logic for the selector */
	case optionsMsg:
		m.selector.options = msg.options
	case moveMsg:
		m.selector.move(msg)
	case tea.KeyMsg:
		switch msg.String() {
		case PluginOptions.Keys.Quit:
			return m, tea.Quit
		case PluginOptions.Keys.Back:
			return m.back(msg)
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

func (m SecondModel) back(msg tea.Msg) (tea.Model, tea.Cmd) {
	firstModel := NewFirstModel()
	return firstModel, firstModel.Init()
}

func (m SecondModel) getOptions() tea.Msg {
	c := &http.Client{Timeout: time.Duration(PluginOptions.Network.Timeout) * time.Millisecond}
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

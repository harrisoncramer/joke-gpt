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

func (m SecondModel) Init() tea.Cmd {
	return m.getOptions
}

func (m SecondModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.err != nil {
		return m, tea.Quit
	}

	/* Handle possible commands by selector */
	updatedSelector, selectorCmd := m.selector.maybeUpdate(msg)
	m.selector = updatedSelector
	if selectorCmd != nil {
		return m, selectorCmd
	}

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
	case tea.KeyMsg:
		switch msg.String() {
		case PluginOptions.Keys.Quit:
			return m, tea.Quit
		case PluginOptions.Keys.Back:
			firstModel := NewFirstModel()
			return firstModel, firstModel.Init()
		}
	}

	return m, cmd
}

func (m SecondModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	base := "Second View\n"
	base += m.selector.View()
	base += fmt.Sprintf("\n\n%s", m.help.View(m.keys))

	return base
}

type OptionsResponse []Option
type optionsMsg struct {
	options []Option
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

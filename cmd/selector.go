package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type selectorOpts struct {
	url     string
	timeout time.Duration
}

type Selector struct {
	opts    selectorOpts
	url     string
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

type selectedEntry struct{}

func newSelector(opts selectorOpts) Selector {

	if opts.timeout == 0 {
		opts.timeout = pluginOpts.Network.Timeout
	}

	m := Selector{
		cursor: 0,
		opts:   opts,
		keys: keyMap{
			Select: selectKeys,
			Up:     upKeys,
			Down:   downKeys,
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
			return selectedEntry{}
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

func (s Selector) getOptions() tea.Msg {
	c := &http.Client{Timeout: s.opts.timeout}
	res, err := c.Get(s.opts.url)

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

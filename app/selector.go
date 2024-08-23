package app

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Options []Option

func (o Options) Filter(search string) []Option {
	if search == "" {
		return o
	}
	var results []Option
	for _, opt := range o {
		if strings.Contains(strings.ToLower(opt.Label), strings.ToLower(search)) {
			results = append(results, opt)
		}
	}
	return results
}

type Selector interface {
	tea.Model
}

type SelectorModel struct {
	cursor  int
	options Options
	filter  textinput.Model
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
	options Options
}

type FilterOpts struct {
	placeholder string
	hidden      bool
}

type NewSelectorModelOpts struct {
	filter  FilterOpts
	options []Option
}

func NewSelectorModel(opts NewSelectorModelOpts) SelectorModel {
	m := SelectorModel{
		options: opts.options,
	}

	if !opts.filter.hidden {
		ti := textinput.New()
		ti.Placeholder = opts.filter.placeholder
		m.filter = ti
	}

	return m
}

func (m SelectorModel) Init() tea.Cmd {
	return nil
}

/* Responds to keypresses that are defined in our plugin options and updates the model, and/or selects a value */
func (m SelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	debugMsg(m, msg)

	var cmd tea.Cmd
	m.filter, cmd = m.filter.Update(msg)
	cmds = append(cmds, cmd)

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
		case PluginOptions.Keys.Filter:
			cmds = append(cmds, textinput.Blink)
			m.filter.Focus()
		case PluginOptions.Keys.Back:
			m.filter.Blur()
		}
	}

	return m, tea.Batch(cmds...)
}

/* Renders the choices and the current cursor */
func (m SelectorModel) View() string {
	base := ""
	base += fmt.Sprintf("%s\n", m.filter.View())
	for i, option := range m.options.Filter(m.filter.Value()) {
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

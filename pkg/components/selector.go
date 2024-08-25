package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/harrisoncramer/joke-gpt/pkg/logger"
	"github.com/harrisoncramer/joke-gpt/shared"
)

type SelectorOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type SelectorOptions []SelectorOption

func (o SelectorOptions) Filter(search string) []SelectorOption {
	if search == "" {
		return o
	}
	var results []SelectorOption
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
	cursor         int
	cursorIcon     string
	options        SelectorOptions
	visibleOptions SelectorOptions
	filter         textinput.Model
	keys           shared.KeyOpts
}

type Direction string

const (
	Up   Direction = "up"
	Down Direction = "down"
)

type FilterOpts struct {
	Placeholder string
	Hidden      bool
}

type NewSelectorModelOpts struct {
	Filter  FilterOpts
	Options []SelectorOption
}

func NewSelectorModel(opts NewSelectorModelOpts) SelectorModel {
	m := SelectorModel{
		options:        opts.Options,
		visibleOptions: opts.Options,
	}

	if !opts.Filter.Hidden {
		ti := textinput.New()
		ti.Placeholder = opts.Filter.Placeholder
		m.filter = ti
	}

	return m
}

func (m SelectorModel) Init() tea.Cmd {
	return nil
}

/* This tea.Msg can be used to set options in a selector */
type SelectorOptionsMsg struct {
	options SelectorOptions
}

/* Responds to keypresses and events, and/or selects a value */
func (m SelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	logger.DebugMsg(m, msg)

	/* Handle our filtering */
	var cmd tea.Cmd
	m.filter, cmd = m.filter.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case SelectorOptionsMsg:
		m.setOptions(msg.options)
	case tea.KeyMsg:
		switch msg.String() {
		case shared.PluginOptions.Keys.Down:
			m.move(Down)
		case shared.PluginOptions.Keys.Up:
			m.move(Up)
		case shared.PluginOptions.Keys.Select:
			return m, m.selectVal
		case shared.PluginOptions.Keys.Filter:
			cmds = append(cmds, textinput.Blink)
			m.filter.Focus()
		case shared.PluginOptions.Keys.Back:
			if len(m.visibleOptions) > 0 {
				m.filter.Blur()
			}
		}
	}

	m.visibleOptions = m.options.Filter(m.filter.Value())
	return m, tea.Batch(cmds...)
}

/* Renders the choices and the current cursor */
func (m SelectorModel) View() string {
	base := ""
	base += fmt.Sprintf("%s\n", m.filter.View())

	if len(m.visibleOptions) == 0 {
		base += "No options found \n"
	} else {
		for i, option := range m.visibleOptions {
			if i == m.cursor {
				base += fmt.Sprintf("%s %s\n", shared.PluginOptions.Display.Cursor, option.Label)
			} else {
				base += fmt.Sprintf("%s %s\n", strings.Repeat(" ", len(m.cursorIcon)), option.Label)
			}
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
		if m.cursor < len(m.visibleOptions)-1 {
			m.cursor++
		}
	}
}

/* Sets options on the selector */
func (m *SelectorModel) setOptions(options []SelectorOption) {
	m.options = options
}

type SelectMsg struct {
	Option SelectorOption
}

/* Chooses the value at the given index */
func (s SelectorModel) selectVal() tea.Msg {
	return SelectMsg{s.options[s.cursor]}
}

package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/harrisoncramer/joke-gpt/pkg/router"
	"github.com/harrisoncramer/joke-gpt/shared"
)

type MultiSelectorOption struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Selected bool   `json:"selected"`
}

type MultiSelectorOptions []MultiSelectorOption

type MultiSelectMsg struct {
	Options MultiSelectorOptions
}

type unselectAllMsg struct{}

type filterOptions struct {
	search string
}

type MultiSelectorModel struct {
	cursor         int
	cursorIcon     string
	options        MultiSelectorOptions
	visibleOptions MultiSelectorOptions
	filter         textinput.Model
	keys           shared.KeyOpts
}

type NewMultiSelectorModelOpts struct {
	Filter  FilterOpts
	Options []MultiSelectorOption
}

func NewMultiSelectorModel(opts NewMultiSelectorModelOpts) MultiSelectorModel {
	m := MultiSelectorModel{
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

func (m MultiSelectorModel) Init() tea.Cmd {
	return func() tea.Msg {
		return unselectAllMsg{}
	}
}

/* Message used to set options in a MultiSelectorModel */
type MultiOptionsMsg struct {
	options MultiSelectorOptions
}

func (m MultiSelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	/* Handle our filtering */
	var cmd tea.Cmd
	m.filter, cmd = m.filter.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case MultiOptionsMsg:
		m.setOptions(msg.options)
	case unselectAllMsg:
		m.unselectAll()
	case tea.KeyMsg:
		switch msg.String() {
		case shared.PluginOptions.Keys.Down:
			m.move(Down)
		case shared.PluginOptions.Keys.Up:
			m.move(Up)
		case shared.PluginOptions.Keys.Select:
			return m, m.confirmSelections
		case shared.PluginOptions.Keys.Toggle:
			m.toggleVal()
		case shared.PluginOptions.Keys.Filter:
			cmds = append(cmds, textinput.Blink)
			m.filter.Focus()
		case shared.PluginOptions.Keys.Back:
			if m.filter.Focused() {
				m.filter.Blur()
				return m, nil
			} else {
				return m, router.Back()
			}
		}
	}

	if m.filter.Focused() {
		m.cursor = 0
	}

	m.visibleOptions = m.options.Filter(m.filterByText)
	return m, tea.Batch(cmds...)
}

func (m MultiSelectorModel) View() string {
	selectedIcon := "x"
	base := ""
	base += fmt.Sprintf("%s\n", m.filter.View())

	if len(m.visibleOptions) == 0 {
		base += "No options found \n"
	} else {
		for i, option := range m.visibleOptions {
			icon := selectedIcon
			if !option.Selected {
				icon = " "
			}
			if i == m.cursor {
				base += fmt.Sprintf("%s [%s] %s\n", shared.PluginOptions.Display.Cursor, icon, option.Label)
			} else {
				base += fmt.Sprintf("%s  [%s] %s\n", strings.Repeat(" ", len(m.cursorIcon)), icon, option.Label)
			}
		}
	}
	return base
}

/* Indicates whether the multi-selector is in a focused state */
func (m MultiSelectorModel) Focused() bool {
	return m.filter.Focused()
}

/* Moves the cursor up or down among the options */
func (m *MultiSelectorModel) move(direction Direction) {
	if m.filter.Focused() {
		return
	}
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

type filterFunc func(o MultiSelectorOption) bool

/* Filter returns a subset of the options based on search and their selected status */
func (o MultiSelectorOptions) Filter(fn filterFunc) MultiSelectorOptions {
	var results []MultiSelectorOption
	for _, opt := range o {
		if fn(opt) {
			results = append(results, opt)
		}
	}
	return results
}

/* Filters the possible options by the text contained in the textinput model */
func (m MultiSelectorModel) filterByText(opt MultiSelectorOption) bool {
	filter := m.filter.Value()
	if filter == "" || strings.Contains(strings.ToLower(opt.Label), strings.ToLower(filter)) {
		return true
	}
	return false
}

/* Filters the possible options by the text contained in the textinput model */
func (m MultiSelectorModel) filterBySelected(opt MultiSelectorOption) bool {
	return opt.Selected
}

func (m *MultiSelectorModel) unselectAll() {
	var results []MultiSelectorOption
	for _, opt := range m.options {
		opt.Selected = false
		results = append(results, opt)
	}
	m.options = results
}

/* Sets options on the selector */
func (m *MultiSelectorModel) setOptions(options []MultiSelectorOption) {
	m.options = options
}

/* Chooses the value at the given index */
func (m *MultiSelectorModel) toggleVal() tea.Msg {
	m.options[m.cursor].Selected = !m.options[m.cursor].Selected
	return nil
}

/* Confirms the selection by sending a multi-select message with all selected options */
func (m MultiSelectorModel) confirmSelections() tea.Msg {
	if m.filter.Focused() {
		return nil
	}

	// TODO: Build "reactive" properties on the model that update with state changes, e.g. selectedOptions
	// For now we can build these options here, and pass them to the message. In the future these computed
	// values should be computed off of m.options
	options := m.options.Filter(m.filterBySelected)
	if len(options) == 0 {
		return nil
	}

	return MultiSelectMsg{options}
}

package router

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Views map[string]tea.Model

type Router struct {
	Model     tea.Model
	Views     Views
	ViewStack []string
	QuitKey   string
}

type changeViewMsg struct {
	view string
}

type backMsg struct{}

type NewRouterModelOpts struct {
	View  string
	Views Views
	Quit  string
}

// Creates a new router that is responsible for handling navigation around the application via the changeView function
func NewRouterModel(opts NewRouterModelOpts) tea.Model {
	r := Router{
		Views:   opts.Views,
		QuitKey: opts.Quit,
	}

	r.setModel(opts.View)
	return r
}

func (m Router) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if cmd := m.handleQuit(msg); cmd != nil {
		return m, cmd
	}

	switch msg := msg.(type) {
	case changeViewMsg:
		m.setModel(msg.view)
		return m, m.Model.Init()
	case backMsg:
		m.back()
		return m, m.Model.Init()
	}

	var cmd tea.Cmd
	m.Model, cmd = m.Model.Update(msg)
	return m, cmd
}

func (m Router) View() string {
	return m.Model.View()
}

func (m Router) Init() tea.Cmd {
	return m.Model.Init()
}

func (m *Router) handleQuit(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case m.QuitKey:
			return tea.Quit
		}
	}
	return nil
}

func (m *Router) setModel(view string) {
	if len(m.ViewStack) == 0 || m.ViewStack[len(m.ViewStack)-1] != view {
		m.Model = m.Views[view]
		m.ViewStack = append(m.ViewStack, view)
	}
}

func (m *Router) back() {
	if len(m.ViewStack) < 2 {
		return
	}
	i := len(m.ViewStack) - 2
	prevView := m.ViewStack[i]
	m.ViewStack = m.ViewStack[:i]
	m.setModel(prevView)
}

// Navigates to the previous router view
func Back() tea.Cmd {
	return func() tea.Msg {
		return backMsg{}
	}
}

// Changes the view of the router
func ChangeView(view string) tea.Cmd {
	return func() tea.Msg {
		return changeViewMsg{view: view}
	}
}

package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/harrisoncramer/joke-gpt/shared"
)

type Router struct {
	Model tea.Model
}

type changeViewMsg struct {
	view string
}

func NewRouterModel(view string) tea.Model {
	m := getModel(view)
	return Router{
		Model: m,
	}
}

// Update method that handles common keystrokes before delegating to the underlying model
func (m Router) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	debugMsg(m, msg)
	if cmd := m.HandleQuit(msg); cmd != nil {
		return m, cmd
	}

	switch msg := msg.(type) {
	case changeViewMsg:
		m.Model = getModel(msg.view)
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

func (m *Router) HandleQuit(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case PluginOptions.Keys.Quit:
			return tea.Quit
		}
	}
	return nil
}

func changeView(view string) tea.Cmd {
	return func() tea.Msg {
		return changeViewMsg{view: view}
	}
}

func getModel(view string) tea.Model {
	switch view {
	case shared.JokeView:
		return NewJokeModel()
	default:
		return NewMainModel()
	}
}

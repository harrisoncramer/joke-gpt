package app

import tea "github.com/charmbracelet/bubbletea"

type NestedView interface {
	tea.Model
	back(msg tea.Msg) (tea.Model, tea.Cmd)
}

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

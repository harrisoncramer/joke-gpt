package cmd

import tea "github.com/charmbracelet/bubbletea"

type Quitter interface {
	tea.Model
	quit(msg tea.Msg) tea.Cmd
}

type NestedView interface {
	tea.Model
	Quitter
	back(msg tea.Msg) Quitter
}

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

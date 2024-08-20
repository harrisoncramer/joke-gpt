package main

import (
	"fmt"
	"slices"

	help "github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	keys keyMap
	help help.Model
}

func newFirstModel() Quitter {
	m := MainModel{
		keys: keyMap{
			Quit:   quitKeys,
			Select: selectKeys,
		},
		help: help.New(),
	}
	return m
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := m.quit(msg)
	if cmd != nil {
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if slices.Contains(m.keys.Select.Keys(), msg.String()) {
			secondModel := newSecondModel()
			return secondModel, secondModel.Init()
		}
	}
	return m, nil
}

func (m MainModel) View() string {
	base := "Main View"
	helpView := m.help.View(m.keys)
	base += fmt.Sprintf("\n\n%s", helpView)
	return base
}

func (m MainModel) quit(msg tea.Msg) tea.Cmd {
	return quit(msg, m.keys.Quit)
}

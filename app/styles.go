package app

import "github.com/charmbracelet/lipgloss"

var textDanger = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FF0000"))

var textGrey = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#585858"))

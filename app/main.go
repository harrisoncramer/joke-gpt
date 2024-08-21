package app

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/harrisoncramer/nested-models/shared"
)

type DisplayOpts struct {
	Cursor string `mapstructure:"cursor"`
}

/* Global plugin options shared across models */
var PluginOptions shared.PluginOpts

/* Initializes the root model and starts the TUI application */
func Start() {
	m := NewFirstModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error starting BubbleTea: %v", err)
		os.Exit(1)
	}
}

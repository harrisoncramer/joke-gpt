package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	cmd "github.com/harrisoncramer/nested-models/cmd"
)

func main() {
	m := cmd.NewFirstModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}

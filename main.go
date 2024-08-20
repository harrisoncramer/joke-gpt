package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TuiOptions struct {
	cursorIcon    string
	globalTimeout time.Duration
}

var tuiOptions = TuiOptions{
	cursorIcon:    ">",
	globalTimeout: time.Second * 2,
}

func main() {
	m := newFirstModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}

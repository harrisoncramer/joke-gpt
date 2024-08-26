package app

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/harrisoncramer/joke-gpt/pkg/router"
	"github.com/harrisoncramer/joke-gpt/shared"
	"github.com/spf13/viper"
)

var appTitle = "GPT Joke 😂\n\n"

/* Initializes the root model and starts the TUI application */
func Start(view string) {
	if viper.GetBool("debug.messages") {
		f, err := tea.LogToFile(shared.PluginOptions.Debug.FilePath, "debug")
		if err != nil {
			fmt.Printf("Error setting up logging: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	m := router.NewRouterModel(router.NewRouterModelOpts{
		Quit: shared.PluginOptions.Keys.Quit,
		Views: map[string]tea.Model{
			shared.JokeView:  NewJokeModel(),
			shared.RootView:  NewMainModel(),
			shared.MultiView: NewMultiChoiceModel(),
		},
		View: view,
	})

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error starting BubbleTea: %v\n", err)
		os.Exit(1)
	}
}

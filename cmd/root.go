package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/spf13/cobra"
)

/* The init() function is called automatically by Go */
func init() {
	rootCmd.PersistentFlags().StringP("token", "t", "", "Token used to authenticate to the Shortcut API")
	rootCmd.PersistentFlags().StringP("config", "", "", "The path to a .yaml configuration file")
	rootCmd.MarkPersistentFlagRequired("token")
}

var rootCmd = &cobra.Command{
	Use:   "sh",
	Short: "A TUI for interacting with Shortcut from the command line",
	Run: func(cmd *cobra.Command, args []string) {
		initializeConfig(cmd)
		m := newFirstModel()
		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			fmt.Printf("There's been an error: %v", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"os"

	app "github.com/harrisoncramer/nested-models/app"
	"github.com/spf13/cobra"
)

/* The init() function is called automatically by Go */
func init() {
	rootCmd.PersistentFlags().StringP("token", "t", "", "Token for the Shortcut API. This value will override a `token` set in your config file")
	rootCmd.PersistentFlags().StringP("config", "", "", "The path to a .yaml configuration file, by default the current directory")
}

var rootCmd = &cobra.Command{
	Use:   "sh",
	Short: "A TUI for interacting with Shortcut from the command line",
	Run: func(cmd *cobra.Command, args []string) {
		err := initializeConfig(cmd)
		if err != nil {
			fmt.Printf("Error configuring application: %v", err)
			os.Exit(1)
		}

		app.Start()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

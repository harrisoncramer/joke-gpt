package cmd

import (
	"fmt"
	"os"

	app "github.com/harrisoncramer/joke-gpt/app"
	"github.com/harrisoncramer/joke-gpt/shared"
	"github.com/spf13/cobra"
)

/* The init() function is called automatically by Go */
func init() {
	rootCmd.PersistentFlags().StringP("token", "t", "", "Token for the ChatGPT API. This value will override a `token` set in your config file. \nIf neither is found, will default to $OPENAI_API_KEY environment variable")
	rootCmd.PersistentFlags().StringP("config", "", "", "The path to a .yaml configuration file, by default the current directory")
}

var rootCmd = &cobra.Command{
	Use:   "joke-gpt",
	Short: "A TUI for interacting with ChatGPT from the command line",
	Run: func(cmd *cobra.Command, args []string) {
		err := initializeConfig(cmd)
		if err != nil {
			fmt.Printf("Error configuring application: %v", err)
			os.Exit(1)
		}

		app.Start(shared.RootView)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

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
	jokeCmd.PersistentFlags().StringP("token", "t", "", "Token for the ChatGPT API. This value will override a `token` set in your config file. \nIf neither is found, will default to $OPEN_API_KEY environment variable")
	jokeCmd.PersistentFlags().StringP("config", "", "", "The path to a .yaml configuration file, by default the current directory")
}

var jokeCmd = &cobra.Command{
	Use:   "joke",
	Short: "Tell a joke immediately",
	Run: func(cmd *cobra.Command, args []string) {
		err := initializeConfig(cmd)
		if err != nil {
			fmt.Printf("Error configuring application: %v", err)
			os.Exit(1)
		}

		app.Start(shared.AppStartArgs{
			Immediate: true,
		})
	},
}

func init() {
	rootCmd.AddCommand(jokeCmd)
}

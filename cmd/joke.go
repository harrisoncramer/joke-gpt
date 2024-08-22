/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	app "github.com/harrisoncramer/joke-gpt/app"
	"github.com/harrisoncramer/joke-gpt/shared"
	"github.com/spf13/cobra"
)

// jokeCmd represents the joke command
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

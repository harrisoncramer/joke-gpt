package cmd

import (
	"os"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type CliOpts struct {
	Network NetworkOpts `yaml:"network"`
	Display DisplayOpts `yaml:"display"`
	Keys    KeyOpts     `yaml:"keys"`
}

type NetworkOpts struct {
	Timeout       int `yaml:"timeout"`
	TimeoutMillis time.Duration
}

type KeyOpts struct {
	Up     string `yaml:"up"`
	Down   string `yaml:"down"`
	Select string `yaml:"enter"`
	Back   string `yaml:"back"`
	Quit   string `yaml:"ctrl+c"`
}

type DisplayOpts struct {
	Cursor string
}

var pluginOpts = CliOpts{}

func initializeConfig(cmd *cobra.Command) error {
	configFile, _ := cmd.Flags().GetString("config")
	if configFile != "" {
		yamlFile, err := os.ReadFile(configFile)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(yamlFile, &pluginOpts)
		if err != nil {
			return err
		}
	}

	if pluginOpts.Display.Cursor == "" {
		pluginOpts.Display.Cursor = ">"
	}
	if pluginOpts.Network.Timeout == 0 {
		pluginOpts.Network.Timeout = 2000
	}
	if pluginOpts.Keys.Up == "" {
		pluginOpts.Keys.Up = "k"
	}
	if pluginOpts.Keys.Down == "" {
		pluginOpts.Keys.Down = "j"
	}
	if pluginOpts.Keys.Select == "" {
		pluginOpts.Keys.Select = "enter"
	}
	if pluginOpts.Keys.Quit == "" {
		pluginOpts.Keys.Quit = "ctrl+c"
	}
	if pluginOpts.Keys.Back == "" {
		pluginOpts.Keys.Back = "esc"
	}

	pluginOpts.Network.TimeoutMillis = time.Duration(pluginOpts.Network.Timeout) * time.Millisecond

	return nil
}

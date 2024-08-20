package cmd

import (
	"os"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type CliOpts struct {
	Debug   bool        `yaml:"debug"`
	Network NetworkOpts `yaml:"network"`
	Display DisplayOpts `yaml:"display"`
}

type NetworkOpts struct {
	Timeout time.Duration `yaml:"timeout"`
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

	defaultOpts := CliOpts{
		Debug: false,
		Network: NetworkOpts{
			Timeout: time.Second * 2,
		},
		Display: DisplayOpts{
			Cursor: ">",
		},
	}

	mergeSettings(defaultOpts)
	return nil
}

func mergeSettings(defaultOpts CliOpts) {
	if pluginOpts.Display.Cursor == "" {
		pluginOpts.Display.Cursor = defaultOpts.Display.Cursor
	}
	if pluginOpts.Network.Timeout == 0 {
		pluginOpts.Network.Timeout = defaultOpts.Network.Timeout
	}
}

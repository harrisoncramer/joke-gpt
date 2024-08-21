package cmd

import (
	"fmt"

	"github.com/harrisoncramer/nested-models/shared"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func initializeConfig(cmd *cobra.Command) (shared.PluginOpts, error) {
	p := shared.PluginOpts{}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetDefault("display.cursor", ">")
	viper.SetDefault("network.timeout", 2000)
	viper.SetDefault("keys.up", "k")
	viper.SetDefault("keys.down", "j")
	viper.SetDefault("keys.select", "enter")
	viper.SetDefault("keys.quit", "ctrl+c")
	viper.SetDefault("keys.back", "esc")

	configFile, _ := cmd.Flags().GetString("config")
	if configFile == "" {
		configFile = "."
	}

	viper.AddConfigPath(configFile)
	err := viper.ReadInConfig()

	if err != nil {
		return p, fmt.Errorf("Fatal error reading configuration file: %v", err)
	}

	flagToken, err := cmd.Flags().GetString("token")
	if err != nil {
		return p, err
	}

	if flagToken != "" {
		viper.Set("token", flagToken)
	}

	if err := viper.Unmarshal(&p); err != nil {
		return p, fmt.Errorf("Fatal error unmarshalling configuration file: %v", err)
	}

	return p, nil
}

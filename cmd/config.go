package cmd

import (
	"fmt"

	app "github.com/harrisoncramer/my-gpt/app"
	"github.com/harrisoncramer/my-gpt/shared"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/* Sets default configuration options then reads in the configuration file and sets it in the app */
func initializeConfig(cmd *cobra.Command) error {
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
	viper.SetDefault("keys.repeat", "r")
	viper.BindPFlag("token", cmd.PersistentFlags().Lookup("token"))

	/* Look for config file in current directory by default */
	configFile, _ := cmd.Flags().GetString("config")
	if configFile == "" {
		configFile = "."
	}
	viper.AddConfigPath(configFile)
	err := viper.ReadInConfig()

	if err != nil {
		return fmt.Errorf("Fatal error reading configuration file: %v", err)
	}

	if err := viper.Unmarshal(&p); err != nil {
		return fmt.Errorf("Fatal error unmarshalling configuration file: %v", err)
	}

	app.PluginOptions = p
	return nil
}

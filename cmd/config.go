package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/harrisoncramer/joke-gpt/shared"
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
	viper.SetDefault("keys.toggle", "t")
	viper.SetDefault("keys.quit", "ctrl+c")
	viper.SetDefault("keys.back", "esc")
	viper.SetDefault("keys.filter", "/")
	viper.SetDefault("keys.repeat", "r")
	viper.SetDefault("keys.help", "?")
	viper.SetDefault("debug.messages", false)
	viper.SetDefault("debug.filepath", "debug.log")
	viper.SetDefault("token", os.Getenv("OPENAI_API_KEY"))
	viper.BindPFlag("token", cmd.PersistentFlags().Lookup("token"))

	/* Look for config file in current directory by default */
	configFile, _ := cmd.Flags().GetString("config")
	if configFile == "" {
		configFile = "."
	}
	viper.AddConfigPath(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("Fatal error reading configuration file: %v\n", err)
		}
	}

	if err := viper.Unmarshal(&p); err != nil {
		return fmt.Errorf("Fatal error unmarshalling configuration file: %v\n", err)
	}

	if viper.Get("token") == "" {
		return errors.New("ChatGPT API Key is required!\n")
	}

	shared.PluginOptions = p
	return nil
}

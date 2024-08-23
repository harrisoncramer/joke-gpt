package shared

/* The options for the plugin, read into the application by Viper from a YAML file */
type PluginOpts struct {
	Token   string      `mapstructure:"token"`
	Network NetworkOpts `mapstructure:"network"`
	Display DisplayOpts `mapstructure:"display"`
	Keys    KeyOpts     `mapstructure:"keys"`
	Debug   DebugOpts   `mapstructure:"debug"`
}

type NetworkOpts struct {
	Timeout int `mapstructure:"timeout"`
}

type KeyOpts struct {
	Up     string `mapstructure:"up"`
	Down   string `mapstructure:"down"`
	Select string `mapstructure:"select"`
	Back   string `mapstructure:"back"`
	Quit   string `mapstructure:"quit"`
	Repeat string `mapstructure:"repeat"`
}

type DisplayOpts struct {
	Cursor string `mapstructure:"cursor"`
}

type DebugOpts struct {
	FilePath    string `mapstructure:"filepath"`
	LogMessages bool   `mapstructure:"log_messages"`
}

type View string

const (
	RootView View = "root"
	JokeView View = "joke"
)

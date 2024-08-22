# joke-gpt

This Repository is a template for creating CLI/TUI tools with Bubble Tea, Cobra, and Viper.

# Configuration

The plugin can use a configuration file in the current directory called `config.yaml` which contains the following schema:

```yaml
token: "api-token"
network:
  timeout: 5000 # In milliseconds
display:
  cursor: ">"
keys:
  up: "k"
  down: "j"
  select: "enter"
  quit: "esc"
  repeat: "r"
```

Run `joke-gpt --help` for more information. You may also supply your ChatGPT token as an environment variable or as a flag argument.

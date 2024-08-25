package logger

import (
	"log"
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
)

func DebugMsg(m tea.Model, msg tea.Msg) {
	if viper.GetBool("debug.messages") {
		modelName := reflect.TypeOf(m)
		log.Printf("Message processed by %s model: %+v", modelName, msg)
	}
}

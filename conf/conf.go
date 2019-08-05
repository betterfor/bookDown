package conf

import (
	"github.com/betterfor/GoLogger/logger"
	"github.com/spf13/viper"
)

const configName = "conf"

func init() {
	viper.AutomaticEnv()
	viper.AddConfigPath("./" + configName)
	viper.SetConfigName(configName)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	initLogs()
}

func initLogs() {
	projectName := viper.GetString("app")
	logger.NewProject("["+projectName+"]", "blue")
	logger.SetLogger(logger.File())
	debug()
}

func debug() {
	logger.AddLogger(logger.ColorConsole())
}

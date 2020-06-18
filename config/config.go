package config

import "github.com/spf13/viper"

var (
	configError error
	configName  = "config"
	configType  = "toml"
	configPath  = "./config"
)

func init() {

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	configError = viper.ReadInConfig()
}

//Error check config init
func Error() error {
	return configError
}

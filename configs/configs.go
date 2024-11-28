// Package configs - Based on Viper
package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

//Configs The mapstructure tags are used to map the struct fields to environment variables
type Configs struct {
	LogLevel string `mapstructure:"LOG_LEVEL"`
	Greeting string `mapstructure:"GREETING"`
}

//LoadEnvViper Loads configuration settings from a file named "app.env" into a Configs struct
func LoadEnvViper() (config Configs)  {
	viper.SetConfigType("env") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".") // optionally look for config in the working directory
	viper.SetConfigFile("app.env")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
	panic(fmt.Errorf("fatal error config file: %s", err))
}
	err = viper.Unmarshal(&config)
	if err != nil {
    panic(err)
}
	return config
}

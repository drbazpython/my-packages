package main

import (
	"drbaz.com/my-packages/configs"
	"drbaz.com/my-packages/logging"
)

//Config ... 
var Config = configs.LoadEnvViper()

//Logger ...
var Logger = logging.DefineLogger(Config.LogLevel)

func main() {
	greeting := Config.Greeting
	Logger.Print("Print Hello World")
	Logger.Printf("Printf %s", greeting)
	Logger.Debug("Debug Hello World")
	Logger.Debugf("Debugf %s", greeting)
	Logger.Info("Info Hello World")
	Logger.Infof("Infof %s", greeting)
	Logger.Error("Error Hello World")
	Logger.Errorf("Errorf %s", greeting)
	Logger.Warn("Warn Hello World")
	Logger.Warnf("Warnf %s", greeting)
	Logger.Fatal("Fatal Hello World")
}


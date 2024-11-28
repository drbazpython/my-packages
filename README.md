## My Packages
 
 ### This repository contains my frequently used packages
 
 ### logging 
 
  - go get github.com/charmbracelet/log
  - go get github.com/charmbracelet/lipgloss
  
  #### Implementation - add these lines to your packages
  
``` go
	import (
		"drbaz.com/my-packages/configs"
		"drbaz.com/my-packages/logging"
		)
	//Config ... 
	var Config = configs.LoadEnvViper()
	//Logger ...
	var Logger = logging.DefineLogger(Config.LogLevel)
```
 ### configs
 
  - go get github.com/spf13/viper
  
  #### Implementation - add these lines to your packages
  
``` go
  import (
    "fmt"
    "github.com/spf13/viper"
    )
    type Configs struct {
      Xxxxxxx string `mapstructure:"XXXXX"` // where XXXXX is defined in your app.env
      Greeting string 'mapstructure:"GREETING"'
      LogLevel string `mapstructure:"LOG_LEVEL"`
    }
```
	
### Makefile

``` 
hello:
	echo "Hello"

build:
	@echo "Building .... "
	@go build -o ./bin drbaz.com/my-packages
	@echo "Installing ...."
	@go install drbaz.com/my-packages
	@echo "my-packages built to bin and installed"
	@echo "Running my-package ...."
	@my-packages
	
run:
	go run main.go

tidy:
	go mod tidy
```

#### main.go

``` go
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
```

### app.env

```
LOG_LEVEL=DEBUG
GREETING="Hello World"
```

### my-go-snippets

~/Library/Application Support/Code/User/snippets/my-go-snippets.json.code-snippets

``` json
{
	".source.go": {
		"varc":
		{
			"prefix":"varc",
			"body":"//Config ... \nvar Config = configs.LoadEnvViper()\n",
			"description":"Snippet for var Config"
		},
		"varlog":
		{
			"prefix":"varl",
			"body":"//Logger ...\nvar Logger = logging.DefineLogger(Config.LogLevel)\n",
			"description":"Snippet for var Logger"
		}
	}
}
```

### CLI packages

#### Bubbletea - The fun, functional and stateful way to build terminal apps

- github.com/charmbraclet/bubbletea

#### Cobra - Cobra is a library providing a simple interface to create powerful modern CLI interfaces

- github.com/spf13/cobra


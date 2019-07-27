package system

import (
	"github.com/github_deployer/src/config"
)

const configPatch = "config.json"

type Application struct {
	Config *config.Config
}

func (app *Application) Init() {
	mainConfig := &config.Config{}
	mainConfig.CreateConfigFile(configPatch)


	mainConfig.ReadConfig(configPatch)
	app.Config = mainConfig
}
package system

import (
	"githubDeployer/src/config"
)

const configPatch = "config.toml"

type Application struct {
	Config *config.Config
}

func (app *Application) Init() {
	config.CreateConfigFile(configPatch)

	mainConfig := &config.Config{}

	if err:= mainConfig.ReadConfig(configPatch); err != nil {
		panic("can`t load config: " + err.Error())
	}
	app.Config = mainConfig
}
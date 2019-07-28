package application

import (
	"encoding/json"
	"fmt"
	"github.com/zeroed88/github-deployer/src/config"
	"github.com/zeroed88/github-deployer/src/logger"
	"github.com/zeroed88/github-deployer/src/services"
	"github.com/zeroed88/github-deployer/src/services/github"
	"log"
	"net/http"
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

	http.HandleFunc("/", app.PushHookHandler)
	fmt.Println("Github Deployer server on 9000")

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


func (app *Application) PushHookHandler(w http.ResponseWriter, r *http.Request) {
	var payload github.PushPayload

	if err := r.ParseForm(); err != nil {
		return
	}

	decoder :=json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		panic(err)
	}

	logMessage := logger.LogMessage{
		Method: r.Method,
		Host: r.Host,
		UrlPath: r.URL.Path,
		Query: r.URL.RawQuery,
	}

	go services.HandlePushPayload(app.Config, payload, logMessage)
}

func New() *Application {
	return &Application{}
}
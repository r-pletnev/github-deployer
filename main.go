package main

import (
	"github.com/github_deployer/src/logger"
	"github.com/github_deployer/src/services"
	"github.com/github_deployer/src/services/github"
	"github.com/github_deployer/src/system"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const logFileName = "log.txt"

func getFile() *os.File {
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	return f
}

func writeLog(logMessage string) {
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)

	log.Println(logMessage)
}

var application = &system.Application{}



func EchoHandler(w http.ResponseWriter, r *http.Request) {
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


	go services.HandlePushPayload(application.Config, payload, logMessage)

}

func main(){
	application.Init()
	http.HandleFunc("/", EchoHandler)
	fmt.Println("Github Deployer server on 9000")

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
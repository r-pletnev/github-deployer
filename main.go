package main

import (
	"githubDeployer/src/services"
	"githubDeployer/src/services/github"
	"githubDeployer/src/system"
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

	services.HandlePushPayload(application.Config, payload)

	logMessage := fmt.Sprintf("method: %v, host: %v, path: %v, query: %v, headers: %v, ",
		r.Method,
		r.Host,
		r.URL.Path,
		r.URL.RawQuery,
		r.Header)

	go writeLog(logMessage)
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
package main

import (
	"fmt"
	"github.com/github_deployer/src/application"
	"github.com/takama/daemon"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	name = "github_deployer"
	description = "Daemon will execute command if github repository will updated"
)

var stdlog, errlog *log.Logger

type Service struct {
	daemon.Daemon
}

// Manage by daemon commands or run the daemon
func (service *Service) Manage() (string, error) {

	usage := "Usage: cron_job install | remove | start | stop | status"
	// If received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			// No need to explicitly stop cron since job will be killed
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	app := &application.Application{}
	app.Init()

	// Waiting for interrupt by system signal
	killSignal := <-interrupt
	stdlog.Println("Got signal:", killSignal)
	return "Service exited", nil
}


func main(){
	srv, err := daemon.New(name, description)
	if err != nil {
		errlog.Println("Error: ", err)
		os.Exit(1)
	}
	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		errlog.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)

}
package services

import (
	"githubDeployer/src/config"
	github2 "githubDeployer/src/services/github"
	"os/exec"
	"runtime"
)

func HandlePushPayload(config *config.Config, payload github2.PushPayload){
	if !payload.IsValidBranch(config.Repository.BranchName){
		return
	}

	go executeCommand(config)
}

func executeCommand(config *config.Config) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/C", config.Repository.Command)
	} else {
		cmd = exec.Command("/bin/sh", config.Repository.Command)
	}

	err := cmd.Run()
	return err
}



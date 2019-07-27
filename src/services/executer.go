package services

import (
	"fmt"
	"github.com/github_deployer/src/config"
	"github.com/github_deployer/src/logger"
	github2 "github.com/github_deployer/src/services/github"
	"os/exec"
	"runtime"
)

func HandlePushPayload(config *config.Config, payload github2.PushPayload, m logger.LogMessage){
	repository := config.GetRepository(payload.Repository.Name)
	var resultErrorText string

	if repository == nil {
		resultErrorText = fmt.Sprintf("repository %s not found", payload.Repository.Name)
		m.WriteToLog(resultErrorText)
		return
	}

	if !payload.IsValidBranch(repository.BranchName){
		resultErrorText = fmt.Sprintf("branch %s not found", payload.BranchName())
		m.WriteToLog(resultErrorText)
		return
	}

	err := executeCommand(repository)
	if err != nil {
		resultErrorText = err.Error()
	}
	m.WriteToLog(resultErrorText)
	return

}

func executeCommand(repository *config.Repository) error {
	var cmd *exec.Cmd


	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/C", repository.Command)
	} else {
		cmd = exec.Command("/bin/sh", repository.Command)
	}

	err := cmd.Run()
	return err
}



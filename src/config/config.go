package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Repository struct {
	Name string `json:"name"`
	BranchName string `json:"branch_name"`
	Event string  `json:"event"`
	Command string `json:"command"`
	Secret string `json:"secret"`
}

type Config struct {
	Repositories []Repository `json:"repositories"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (c *Config) ReadConfig(configPath string) {
	body, err := ioutil.ReadFile(configPath)
	check(err)

	err = json.Unmarshal(body, c)
	check(err)
}


func (c *Config) CreateConfigFile(path string) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return
	}

	exampleRepository := Repository{
		"Repository Name",
		"master",
		"push",
		"./command.sh",
		"",
	}

	c.Repositories = append(c.Repositories, exampleRepository)

	body, err := json.MarshalIndent(c, "", "  ")
	check(err)
	err = ioutil.WriteFile(path, body, 0655)
	check(err)
}

func (c *Config) GetRepository(repositoryName string) *Repository {
	for _, repo := range c.Repositories {
		if repo.Name == repositoryName {
			return &repo
		}
	}

	return nil
}




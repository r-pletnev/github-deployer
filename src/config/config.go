package config

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"os"
)

type GeneralConfig struct {
	BranchName string `toml:"branch_name"`
	Event string  `toml:"event"`
	Command string `toml:"command"`
	Secret string `toml:"secret"`
}

type Config struct {
	Repository GeneralConfig `toml:"general"`
}

func (c *Config) ReadConfig(configPath string) error {
	tree, err := toml.LoadFile(configPath)
	if err != nil {
		return err
	}
	return tree.Unmarshal(c)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateConfigFile(patch string) {
	if _, err := os.Stat(patch); !os.IsNotExist(err) {
		return
	}

	body := []byte("[general]\nbranch_name = \"master\"\nevent = \"push\"\ncommand = \"\"\nsecret = \"\"\n")
	err := ioutil.WriteFile(patch, body, 0655)
	check(err)
}

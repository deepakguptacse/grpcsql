package configs

import (
	"fmt"
	"io/ioutil"

	"github.com/deepakguptacse/grpcsql/env"
	"gopkg.in/yaml.v2"
)

var envToConfigFile = map[env.Environment]string{
	env.Dev:  "configs/dev.yaml",
	env.Prod: "configs/prod.yaml",
}

type Config struct {
	SQLAddress string `yaml:"sql_address"`
}

func ReadConfig() (*Config, error) {
	var cfg Config

	var configFilePath string
	configFilePath, ok := envToConfigFile[env.GetCurrentEnv()]
	if !ok {
		return nil, fmt.Errorf("no config file found for the current environment")
	}

	yamlFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("couldn't read file: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return nil, fmt.Errorf("couldn't unmarshal file: %v", err)
	}
	return &cfg, nil
}

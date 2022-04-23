package config

import (
	"fmt"
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// ENV for environment
var ENV string = "dev"

type Config struct {
	LogLevel string `yaml:"log_level",envconfig:"LOG_LEVEL"`
	Server   struct {
		Port int    `yaml:"port", envconfig:"SERVER_PORT"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Host           string `yaml:"host", envconfig:"DB_HOST"`
		Port           int    `yaml:"port", envconfig:"DB_PORT"`
		Username       string `yaml:"user", envconfig:"DB_USERNAME"`
		Password       string `yaml:"pass",envconfig:"DB_PASSWORD"`
		MaxConnections int    `yaml:"max_connections"`
	} `yaml:"database"`
}

func Load(env string) (*Config, error) {

	var config Config
	var err error
	var file []byte

	err = envconfig.Process("", &config)
	if err != nil {
		return &config, err
	}

	if file, err = ioutil.ReadFile(fmt.Sprintf("config/config.%s.yaml", env)); err == nil {
		if err = yaml.Unmarshal(file, &config); err == nil {
			return &config, nil
		}

		return nil, fmt.Errorf("config file is malformed %v", err)
	}
	return &config, fmt.Errorf("config file is non existent %v", err)

}

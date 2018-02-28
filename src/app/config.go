package app

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type dbConfig struct {
	Name    string
	User    string
	Pass    string
	Host    string
	Port    int
	SslMode bool
}

type appConfig struct {
	Port int
}

type yamlConfig struct {
	Db  dbConfig
	App appConfig
}

// Config stores application configuration
var Config yamlConfig

func LoadConfig() {
	source, err := ioutil.ReadFile("config/config.yaml")

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(source, &Config)

	if err != nil {
		panic(err)
	}
}

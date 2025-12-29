package main

import (
	"fmt"

	"github.com/pfremaux/golibs/config/pkg/config"
)

type AppConfig struct {
	DatabaseURL string `yaml:"database_url"`
	Port        int    `yaml:"port"`
}

func main() {
	configFromFlag := config.LoadFlagsConfig([]config.Parameter{
		{
			Key:         "config",
			DefaultVal:  "config.yaml",
			Description: "Config file path",
		},
	})
	cfg := AppConfig{}
	err := config.LoadYaml(*configFromFlag["config"], &cfg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", cfg.DatabaseURL)
}

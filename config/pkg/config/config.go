package config

import (
	"flag"
	"os"

	"github.com/pfremaux/golibs/files/pkg/files"
	"gopkg.in/yaml.v3"
)

func GetUserHomeDir() string {
	if home, err := os.UserHomeDir(); err == nil {
		return home
	}
	return ""
}

func InitConfigDir(appName string) string {
	homeDir := GetUserHomeDir()
	configDir := homeDir + string(os.PathSeparator) + ".config" + string(os.PathSeparator) + appName
	files.MkDirIfNotExists(configDir)
	return configDir
}

type Parameter struct {
	Key         string
	DefaultVal  string
	Description string
}

func LoadFlagsConfig(params []Parameter) map[string]*string {
	configs := make(map[string]*string)

	for _, param := range params {
		val := flag.String(param.Key, param.DefaultVal, param.Description)
		configs[param.Key] = val
	}
	flag.Parse()
	return configs
}

func LoadYaml(filePath string, out any) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, out); err != nil {
		return err
	}
	return nil
}

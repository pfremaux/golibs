package config

import "os"

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

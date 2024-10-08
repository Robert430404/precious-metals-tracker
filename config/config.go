package config

import (
	"os"
)

type Config struct {
	ConfigPath string
}

var HydratedConfig *Config = nil

func (self *Config) Hydrate() {
	var homeDir string = os.Getenv("HOME")
	if len(homeDir) == 0 {
		panic("could not load home home directory")
	}

	self.ConfigPath = homeDir + "/.local/share/precious-metals-tracker"
}

func GetConfig() *Config {
	if HydratedConfig != nil {
		return HydratedConfig
	}

	HydratedConfig := &Config{}
	HydratedConfig.Hydrate()

	return HydratedConfig
}
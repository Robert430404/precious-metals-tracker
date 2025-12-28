package config

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	ConfigPath   string
	SqlitePath   string
	RuntimeFlags Flags
}

var HydratedConfig *Config = nil

func (self *Config) Hydrate() error {
	var homeDir string = os.Getenv("HOME")
	if len(homeDir) == 0 {
		return errors.New("could not load home home directory")
	}

	self.ConfigPath = homeDir + "/.local/share/precious-metals-tracker"

	var overrideConfigPath = os.Getenv("PRECIOUS_METALS_TRACKER_DATA_DIR")
	if len(overrideConfigPath) > 0 {
		self.ConfigPath = overrideConfigPath
	}

	flags := Flags{}
	flags.Hydrate(self.ConfigPath)

	self.RuntimeFlags = flags

	self.SqlitePath = self.ConfigPath + "/precious-metals-tracker.sqlite"

	return nil
}

func GetConfig() (*Config, error) {
	if HydratedConfig != nil {
		return HydratedConfig, nil
	}

	HydratedConfig := &Config{}
	err := HydratedConfig.Hydrate()
	if err != nil {
		return nil, fmt.Errorf("there was a problem hydrating your configuration: %v", err)
	}

	return HydratedConfig, nil
}

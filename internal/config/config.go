package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

type Config struct {
	BotId     string   `yaml:"bot_id"`
	Paths     []string `yaml:"paths"`
	TargetDir string   `yaml:"target_dir"`
}

var Cfg *Config

func Load() error {
	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed: could not get current os user: %s", err.Error())
	}

	configPath := filepath.Join(currentUser.HomeDir, "btool.yaml")
	file, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to open config file: %s", err.Error())
	}

	c := &Config{}
	if err := yaml.Unmarshal(file, c); err != nil {
		return fmt.Errorf("failed to read config file: %s", err.Error())
	}

	Cfg = c

	return nil
}

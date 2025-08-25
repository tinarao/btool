package config

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

type Config struct {
	BotToken       string   `yaml:"bot_token"`
	Paths          []string `yaml:"paths"`
	TargetDir      string   `yaml:"target_dir"`
	ChatId         int64    `yaml:"chat_id"`
	LastBackupDate string   `yaml:"last_backup_date,omitempty"`
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

func (c *Config) SetLastBackupTime(isoDate string) {
	// TODO backup and revert if err != nil
	c.LastBackupDate = isoDate
	c.save()
}

func GetConfigPath() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("failed: could not get current os user: %s", err.Error())
	}

	return filepath.Join(currentUser.HomeDir, "btool.yaml"), nil
}

func (c *Config) save() error {
	serialized, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	err = os.WriteFile(path, serialized, 0644)
	return err
}

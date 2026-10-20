package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	Db_url				string
	Current_user_name	string
}

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	readConfig, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, fmt.Errorf("Error on reading gator config file: %s", err)
	}

	var config Config
	if err := json.Unmarshal(readConfig, &config); err != nil {
		return Config{}, fmt.Errorf("Error on decoding config json: %s", err)
	}

	return config, nil
}

func (cfg *Config) SetUser(user string) error {
	cfg.Current_user_name = user
	return write(*cfg)
}

func write(cfg Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	cfgData, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Error encoding json data: %s", err)
	}

	err = os.WriteFile(configFilePath, cfgData, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to config file: %s", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Couldn't get home directory %s", err)
	}
	gatorConfigDir := homeDir + configFileName

	return gatorConfigDir, nil
}
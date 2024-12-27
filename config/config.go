package config

import (
	"fmt"
	"io"
	"kita-go-scheduler/constants"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Tasks []TaskConfig `mapstructure:"tasks"`
	Logs  LogConfig    `mapstructure:"logs"`
}

type TaskConfig struct {
	Name     string `mapstructure:"name"`
	Schedule string `mapstructure:"schedule"`
	Command  string `mapstructure:"command"`
}

type LogConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// LoadConfig loads the configuration from ~/.config/{constants.PROJECT_NAME}/config.yaml
func LoadConfig() Config {
	// Get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get home directory: %v", err)
	}

	// Build the path to the config file
	configDir := filepath.Join(homeDir, ".config", constants.PROJECT_NAME)
	configFile := filepath.Join(configDir, constants.CONFIG_FILE)

	// Check if the config directory exists, if not create it
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err = os.MkdirAll(configDir, 0755)
		if err != nil {
			log.Fatalf("Failed to create config directory: %v", err)
		}
		log.Printf("Created config directory: %s", configDir)
	}

	// Check if the config file exists, if not copy the example file
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		exampleConfigPath, err := getExampleConfigPath()
		if err != nil {
			log.Fatalf("Failed to determine example config file path: %v", err)
		}

		if err := copyFile(exampleConfigPath, configFile); err != nil {
			log.Fatalf("Failed to copy example config file: %v", err)
		}
		log.Printf("Copied example %v to: %s", constants.CONFIG_FILE, configFile)
	}

	// Set up viper to read the config file
	viper.SetConfigFile(configFile)
	viper.SetConfigType(constants.CONFIG_FILE_EXT)

	// Create an empty config struct
	var config Config

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Unmarshal the config into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	return config
}

// getExampleConfigPath determines the absolute path to the example config file
func getExampleConfigPath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %v", err)
	}
	execDir := filepath.Dir(execPath)
	exampleConfigPath := filepath.Join(execDir, constants.CONFIG_FILE)

	// Check if the example file exists
	if _, err := os.Stat(exampleConfigPath); os.IsNotExist(err) {
		return "", fmt.Errorf("example config file not found at %s", exampleConfigPath)
	}

	return exampleConfigPath, nil
}

// copyFile copies a file from src to dst.
func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %v", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %v", err)
	}

	return nil
}

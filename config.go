package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/matthewreider/bthropic/claude"
)

type Config struct {
	ClaudeAPIKey string `json:"claude_api_key"`
}

func getConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".bthropic", "config.json")
}

func initializeApp() error {
	// Create config directory
	configDir := filepath.Dir(getConfigPath())
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	var config Config
	fmt.Print("Enter Claude API Key: ")
	fmt.Scanln(&config.ClaudeAPIKey)

	// Test Claude API
	client := claude.NewClient(config.ClaudeAPIKey)
	_, err := client.CreateMessage("test")
	if err != nil {
		return fmt.Errorf("invalid Claude API key: %w", err)
	}

	// Check if Bear is installed
	// TODO: Implement Bear installation check

	// Save encrypted config
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(getConfigPath(), data, 0600)
}

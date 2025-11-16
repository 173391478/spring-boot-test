
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Config represents application configuration
type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
	Features FeatureFlags   `json:"features"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DatabaseConfig struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type FeatureFlags struct {
	EnableCache bool `json:"enable_cache"`
	DebugMode   bool `json:"debug_mode"`
}

// ParseConfig parses JSON configuration with comprehensive error handling
func ParseConfig(jsonData []byte) (*Config, error) {
	var config Config
	
	if err := json.Unmarshal(jsonData, &config); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}
	
	if err := validateConfig(&config); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}
	
	return &config, nil
}

// validateConfig performs basic validation on configuration
func validateConfig(config *Config) error {
	if config.Server.Host == "" {
		return fmt.Errorf("server host cannot be empty")
	}
	
	if config.Server.Port <= 0 || config.Server.Port > 65535 {
		return fmt.Errorf("invalid server port: %d", config.Server.Port)
	}
	
	if config.Database.URL == "" {
		return fmt.Errorf("database URL cannot be empty")
	}
	
	return nil
}

// ConfigToString converts config to formatted JSON string
func (c *Config) String() string {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error formatting config: %v", err)
	}
	return string(bytes)
}

func main() {
	// Example usage
	jsonConfig := `{
		"server": {
			"host": "localhost",
			"port": 8080
		},
		"database": {
			"url": "postgres://localhost:5432/mydb",
			"username": "admin",
			"password": "secret"
		},
		"features": {
			"enable_cache": true,
			"debug_mode": false
		}
	}`
	
	config, err := ParseConfig([]byte(jsonConfig))
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}
	
	fmt.Println("Configuration parsed successfully:")
	fmt.Println(config.String())
	
	// Demonstrate accessing configuration values
	fmt.Printf("\nServer will run on: %s:%d\n", config.Server.Host, config.Server.Port)
	fmt.Printf("Database features - Cache: %t, Debug: %t\n", 
		config.Features.EnableCache, config.Features.DebugMode)
}
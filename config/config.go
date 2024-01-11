// Package config provides the functionality to interact with the configuration of the application.
package config

import (
	"errors"                 // Errors package provides the functionality to create error messages.
	"github.com/spf13/viper" // Viper is a complete configuration solution for Go applications.
	"log"                    // Log package provides the functionality to implement logging.
)

// Config struct represents the configuration of the application with fields for the server and database configurations.
// Server: The server configuration of the application.
// DB: The database configuration of the application.
type Config struct {
	Server ServerConfig   `mapstructure:"app"` // The server configuration of the application.
	DB     DatabaseConfig `mapstructure:"db"`  // The database configuration of the application.
}

// ServerConfig struct represents the server configuration with fields for the host, port, mode, and debug.
// Host: The host of the server.
// Port: The port of the server.
// Mode: The mode of the server.
// Debug: The debug mode of the server.
type ServerConfig struct {
	Host  string `mapstructure:"host"`  // The host of the server.
	Port  int    `mapstructure:"port"`  // The port of the server.
	Mode  string `mapstructure:"mode"`  // The mode of the server.
	Debug bool   `mapstructure:"debug"` // The debug mode of the server.
}

// DatabaseConfig struct represents the database configuration with fields for the database type and SQLite configuration.
// DatabaseType: The type of the database.
// Sqlite: The SQLite configuration.
type DatabaseConfig struct {
	DatabaseType string       `mapstructure:"database_type"` // The type of the database.
	Sqlite       SqliteConfig `mapstructure:"sqlite"`        // The SQLite configuration.
}

// SqliteConfig struct represents the SQLite configuration with a field for the database path.
// DatabasePath: The path of the SQLite database.
type SqliteConfig struct {
	DatabasePath string `mapstructure:"database_path"` // The path of the SQLite database.
}

// NewConfig creates a new configuration by reading from a YAML file and environment variables.
// It uses Viper to read the configuration.
// If the configuration file is not found, it returns an error.
// If the configuration file is found, it unmarshals the configuration into a Config object and returns it.
// Returns a Config object and an error.
func NewConfig() (*Config, error) {
	v := viper.New()                 // Creates a new Viper instance.
	v.SetConfigType("yaml")          // Sets the configuration type to YAML.
	v.SetConfigName("config/config") // Sets the configuration file name.
	v.AddConfigPath(".")             // Adds the current directory as a path to look for the configuration file.
	v.AutomaticEnv()                 // Reads in environment variables that match.

	// Reads the configuration file.
	// If the configuration file is not found, it returns an error.
	// If the configuration file is found, it unmarshals the configuration into a Config object.
	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	cfg := &Config{}                         // Creates a new Config object.
	if err := v.Unmarshal(cfg); err != nil { // Unmarshals the configuration into the Config object.
		return nil, err
	}

	log.Printf("Config loaded from %v", v.AllSettings()) // Logs the loaded configuration.

	return cfg, nil // Returns the Config object.
}

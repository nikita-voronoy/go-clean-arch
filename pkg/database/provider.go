// Package database provides the functionality to create a new database connection.
package database

import (
	"errors"
	"x-ci-cd/config"
	"x-ci-cd/pkg/database/sqlite"
)

// NewDatabase creates a new database connection based on the provided configuration.
// It currently supports only SQLite databases.
// cfg: The configuration object that contains the database settings.
// Returns a Database object if the database connection is successfully established.
// Returns an error if the database type is not supported or if the connection cannot be established.
func NewDatabase(cfg *config.Config) (Database, error) {
	switch cfg.DB.DatabaseType {
	case "sqlite":
		// Create a new SQLite database connection.
		return sqlite.NewDatabase(cfg)
	default:
		// Return an error if the database type is not supported.
		return nil, errors.New("database type not supported")
	}
}

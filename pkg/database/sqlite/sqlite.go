// Package sqlite provides the functionality to interact with a SQLite database.
package sqlite

import (
	"context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"x-ci-cd/config"
	"x-ci-cd/internal/entities"
)

// Database struct represents a SQLite database connection.
type Database struct {
	db *gorm.DB
}

// NewDatabase creates a new SQLite database connection based on the provided configuration.
// cfg: The configuration object that contains the SQLite database settings.
// Returns a Database object if the database connection is successfully established.
// Returns an error if the connection cannot be established or if the database migration fails.
func NewDatabase(cfg *config.Config) (*Database, error) {
	conn, err := gorm.Open(sqlite.Open(cfg.DB.Sqlite.DatabasePath), &gorm.Config{})
	if err != nil {
		return nil, err // return an error instead of panicking
	}
	if err := conn.AutoMigrate(entities.UserLogin{}, entities.User{Metadata: entities.Metadata{}}); err != nil {
		return nil, err
	}
	return &Database{db: conn}, nil
}

// Create adds a new record to the SQLite database.
// ctx: The context for the operation.
// entity: The record to add.
// Returns an error if the operation fails.
func (g Database) Create(ctx context.Context, entity interface{}) error {
	return g.db.WithContext(ctx).Create(entity).Error
}

// Read retrieves a record from the SQLite database.
// ctx: The context for the operation.
// entity: The record to retrieve.
// compareString: The field to compare.
// compareValues: The values to compare.
// Returns an error if the operation fails.
func (g Database) Read(ctx context.Context, entity interface{}, compareString string, compareValues ...interface{}) error {
	return g.db.WithContext(ctx).Where(compareString, compareValues...).First(entity).Error
}

// Update modifies a record in the SQLite database.
// ctx: The context for the operation.
// entity: The record to modify.
// Returns an error if the operation fails.
func (g Database) Update(ctx context.Context, entity interface{}) error {
	return g.db.WithContext(ctx).Save(entity).Error
}

// Delete removes a record from the SQLite database.
// ctx: The context for the operation.
// entity: The record to remove.
// id: The id of the record to remove.
// Returns an error if the operation fails.
func (g Database) Delete(ctx context.Context, entity interface{}, id interface{}) error {
	return g.db.WithContext(ctx).Where("id = ?", id).Delete(entity).Error
}

// ReadAll retrieves all records from the SQLite database.
// ctx: The context for the operation.
// entity: The records to retrieve.
// Returns an error if the operation fails.
func (g Database) ReadAll(ctx context.Context, entity interface{}) error {
	result := g.db.WithContext(ctx).Find(entity)
	return result.Error
}

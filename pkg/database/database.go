// Package database provides the functionality to interact with a database.
package database

import (
	"golang.org/x/net/context"
)

// Database is an interface that defines the methods required for database operations.
// It includes methods for creating, reading, updating, and deleting records.
// Each method requires a context and an entity.
// The entity is the record that needs to be created, read, updated, or deleted.
// For the Read method, a compareString and compareValue are also required to find the record.
// For the Delete method, an id is required to find the record.
type Database interface {
	// Create adds a new record to the database.
	// ctx: The context for the operation.
	// entity: The record to add.
	// Returns an error if the operation fails.
	Create(ctx context.Context, entity interface{}) error

	// Read retrieves a record from the database.
	// ctx: The context for the operation.
	// entity: The record to retrieve.
	// compareString: The field to compare.
	// compareValue: The value to compare.
	// Returns an error if the operation fails.
	Read(ctx context.Context, entity interface{}, compareString string, compareValue ...interface{}) error

	// Update modifies a record in the database.
	// ctx: The context for the operation.
	// entity: The record to modify.
	// Returns an error if the operation fails.
	Update(ctx context.Context, entity interface{}) error

	// Delete removes a record from the database.
	// ctx: The context for the operation.
	// entity: The record to remove.
	// id: The id of the record to remove.
	// Returns an error if the operation fails.
	Delete(ctx context.Context, entity interface{}, id interface{}) error

	// ReadAll retrieves all records from the database.
	// ctx: The context for the operation.
	// entity: The records to retrieve.
	// Returns an error if the operation fails.
	ReadAll(ctx context.Context, entity interface{}) error
}

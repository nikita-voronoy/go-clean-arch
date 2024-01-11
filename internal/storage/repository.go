// Package storage provides the functionality to interact with user data in the storage.
package storage

import (
	"context"
	"github.com/google/uuid"
	"x-ci-cd/internal/entities"
)

// UserRepository is an interface that defines the methods required for user data operations.
// It includes methods for creating, reading, updating, and deleting user records.
// Each method requires a context and an entity.
// The entity is the user record that needs to be created, read, updated, or deleted.
// For the Read method, an id is required to find the user record.
// For the ReadByEmail and ReadByUsername methods, an email and a username are required respectively to find the user record.
// The CheckUserExists method requires an email and a username to check if the user exists.
type UserRepository interface {
	// Create adds a new user record to the storage.
	// ctx: The context for the operation.
	// model: The user record to add.
	// Returns an error if the operation fails.
	Create(ctx context.Context, model entities.User) error

	// Read retrieves a user record from the storage.
	// ctx: The context for the operation.
	// id: The id of the user record to retrieve.
	// Returns the user record and an error if the operation fails.
	Read(ctx context.Context, id uuid.UUID) (entities.User, error)

	// Update modifies a user record in the storage.
	// ctx: The context for the operation.
	// model: The user record to modify.
	// Returns an error if the operation fails.
	Update(ctx context.Context, model entities.User) error

	// Delete removes a user record from the storage.
	// ctx: The context for the operation.
	// id: The id of the user record to remove.
	// Returns an error if the operation fails.
	Delete(ctx context.Context, id uuid.UUID) error

	// ReadByEmail retrieves a user record from the storage based on the email.
	// ctx: The context for the operation.
	// email: The email of the user record to retrieve.
	// Returns the user record and an error if the operation fails.
	ReadByEmail(ctx context.Context, email string) (entities.User, error)

	// ReadByUsername retrieves a user record from the storage based on the username.
	// ctx: The context for the operation.
	// username: The username of the user record to retrieve.
	// Returns the user record and an error if the operation fails.
	ReadByUsername(ctx context.Context, username string) (entities.User, error)

	// ReadAll retrieves all user records from the storage.
	// ctx: The context for the operation.
	// model: The user records to retrieve.
	// Returns the user records and an error if the operation fails.
	ReadAll(ctx context.Context, model []entities.User) ([]entities.User, error)

	// CheckUserExists checks if a user exists in the storage based on the email and username.
	// ctx: The context for the operation.
	// email: The email of the user to check.
	// username: The username of the user to check.
	// Returns a boolean indicating if the user exists and an error if the operation fails.
	CheckUserExists(ctx context.Context, email string, username string) (bool, error)
}

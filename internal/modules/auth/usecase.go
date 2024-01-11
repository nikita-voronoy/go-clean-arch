// Package auth provides the functionality to interact with user authentication data.
package auth

import (
	"context"
	"github.com/google/uuid"
	"github.com/nikita-voronoy/go-clean-arch/internal/entities"
)

// UseCase is an interface that defines the methods required for user authentication operations.
// It includes methods for registering, logging in, getting all users, hashing and comparing passwords, generating UUIDs and bearer tokens, and validating users.
// Each method requires a context and an entity.
// The entity is the user or user login record that needs to be processed.
type UseCase interface {
	// Register adds a new user record to the storage.
	// ctx: The context for the operation.
	// user: The user record to add.
	// Returns an error if the operation fails.
	Register(ctx context.Context, user entities.User) error

	// Login checks the user credentials and logs in the user.
	// ctx: The context for the operation.
	// user: The user login record to check.
	// Returns a string and an error if the operation fails.
	Login(ctx context.Context, user entities.UserLogin) (string, error)

	// GetAll retrieves all user records from the storage.
	// ctx: The context for the operation.
	// Returns the user records and an error if the operation fails.
	GetAll(ctx context.Context) ([]entities.User, error)

	// HashPassword hashes the provided password.
	// password: The password to hash.
	// Returns the hashed password and an error if the operation fails.
	HashPassword(password string) (string, error)

	// ComparePasswords compares the hashed password and the provided password.
	// hashedPassword: The hashed password to compare.
	// password: The password to compare.
	// Returns an error if the passwords do not match.
	ComparePasswords(hashedPassword, password string) error

	// GenerateUUID generates a new UUID.
	// Returns the generated UUID and an error if the operation fails.
	GenerateUUID() (uuid.UUID, error)

	// GenerateBearerToken generates a new bearer token.
	// Returns the generated bearer token and an error if the operation fails.
	GenerateBearerToken() (string, error)

	// Validate validates the provided user record.
	// user: The user record to validate.
	// Returns an error if the user record is not valid.
	Validate(user entities.User) error
}

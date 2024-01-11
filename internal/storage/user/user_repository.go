// Package user provides the functionality to interact with user data in the storage.
package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/nikita-voronoy/go-clean-arch/internal/entities"
	"github.com/nikita-voronoy/go-clean-arch/internal/storage"
	"github.com/nikita-voronoy/go-clean-arch/pkg/database"
)

// Repository struct represents a user repository that provides methods for user data operations.
type Repository struct {
	db database.Database
}

// Create adds a new user record to the storage.
// ctx: The context for the operation.
// model: The user record to add.
// Returns an error if the operation fails.
func (r Repository) Create(ctx context.Context, model entities.User) error {
	if err := r.db.Create(ctx, &model); err != nil {
		return err
	}
	return nil
}

// Read retrieves a user record from the storage.
// ctx: The context for the operation.
// id: The id of the user record to retrieve.
// Returns the user record and an error if the operation fails.
func (r Repository) Read(ctx context.Context, id uuid.UUID) (entities.User, error) {
	var user entities.User
	result := r.db.Read(ctx, &user, "id = ?", id)
	if result != nil {
		return entities.User{}, result
	}
	return user, nil
}

// Update modifies a user record in the storage.
// ctx: The context for the operation.
// model: The user record to modify.
// Returns an error if the operation fails.
func (r Repository) Update(ctx context.Context, model entities.User) error {
	if err := r.db.Update(ctx, &model); err != nil {
		return err
	}
	return nil
}

// Delete removes a user record from the storage.
// ctx: The context for the operation.
// id: The id of the user record to remove.
// Returns an error if the operation fails.
func (r Repository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.Delete(ctx, entities.User{}, id); err != nil {
		return err
	}
	return nil
}

// ReadByEmail retrieves a user record from the storage based on the email.
// ctx: The context for the operation.
// email: The email of the user record to retrieve.
// Returns the user record and an error if the operation fails.
func (r Repository) ReadByEmail(ctx context.Context, email string) (entities.User, error) {
	var user entities.User
	if err := r.db.Read(ctx, &user, "email = ?", email); err != nil {
		return entities.User{}, errors.New("record not found")
	}
	return user, nil
}

// ReadByUsername retrieves a user record from the storage based on the username.
// ctx: The context for the operation.
// username: The username of the user record to retrieve.
// Returns the user record and an error if the operation fails.
func (r Repository) ReadByUsername(ctx context.Context, username string) (entities.User, error) {
	var user entities.User
	if err := r.db.Read(ctx, &user, "username = ?", username); err != nil {
		return entities.User{}, errors.New("record not found")
	}
	return user, nil
}

// CheckUserExists checks if a user exists in the storage based on the email and username.
// ctx: The context for the operation.
// email: The email of the user to check.
// username: The username of the user to check.
// Returns a boolean indicating if the user exists and an error if the operation fails.
func (r Repository) CheckUserExists(ctx context.Context, email string, username string) (bool, error) {
	var user entities.User
	err := r.db.Read(ctx, &user, "email = ? OR username = ?", email, username)
	if err != nil {
		if errors.Is(err, errors.New("record not found")) {
			// If no record is found, it means the user doesn't exist.
			// Return false and no error.
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// ReadAll retrieves all user records from the storage.
// ctx: The context for the operation.
// model: The user records to retrieve.
// Returns the user records and an error if the operation fails.
func (r Repository) ReadAll(ctx context.Context, model []entities.User) ([]entities.User, error) {
	if err := r.db.ReadAll(ctx, &model); err != nil {
		return nil, err
	}
	return model, nil
}

// NewUserRepository creates a new user repository with the provided database.
// db: The database for the user repository.
// Returns a UserRepository object.
func NewUserRepository(db database.Database) storage.UserRepository {
	return &Repository{
		db: db,
	}
}

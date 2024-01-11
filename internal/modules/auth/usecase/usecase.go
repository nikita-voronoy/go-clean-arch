// Package usecase provides the functionality to interact with user authentication data.
package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/nikita-voronoy/go-clean-arch/config"
	"github.com/nikita-voronoy/go-clean-arch/internal/entities"
	"github.com/nikita-voronoy/go-clean-arch/internal/modules/auth"
	"github.com/nikita-voronoy/go-clean-arch/internal/storage"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// AuthUseCase struct represents a user authentication use case that provides methods for user authentication operations.
type AuthUseCase struct {
	cfg  *config.Config
	repo storage.UserRepository
}

// NewAuthUC creates a new user authentication use case with the provided configuration and user repository.
// cfg: The configuration for the user authentication use case.
// repo: The user repository for the user authentication use case.
// Returns an auth.UseCase object.
func NewAuthUC(cfg *config.Config, repo storage.UserRepository) auth.UseCase {
	return &AuthUseCase{
		cfg:  cfg,
		repo: repo,
	}
}

// Validate validates the provided user record.
// user: The user record to validate.
// Returns an error if the user record is not valid.
func (uc AuthUseCase) Validate(user entities.User) error {
	return validator.New().Struct(user)
}

// HashPassword hashes the provided password.
// password: The password to hash.
// Returns the hashed password and an error if the operation fails.
func (uc AuthUseCase) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}
	return string(hashedPassword), nil
}

// ComparePasswords compares the hashed password and the provided password.
// hashedPassword: The hashed password to compare.
// password: The password to compare.
// Returns an error if the passwords do not match.
func (uc AuthUseCase) ComparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// GenerateUUID generates a new UUID.
// Returns the generated UUID and an error if the operation fails.
func (uc AuthUseCase) GenerateUUID() (uuid.UUID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return uuid.UUID{}, err
	}
	return id, nil
}

// GenerateBearerToken generates a new bearer token.
// Returns the generated bearer token and an error if the operation fails.
func (uc AuthUseCase) GenerateBearerToken() (string, error) {
	token := make([]byte, 128)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(token), nil
}

// GetAll retrieves all user records from the storage.
// ctx: The context for the operation.
// Returns the user records and an error if the operation fails.
func (uc AuthUseCase) GetAll(ctx context.Context) ([]entities.User, error) {
	var users []entities.User
	users, err := uc.repo.ReadAll(ctx, users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Register adds a new user record to the storage.
// ctx: The context for the operation.
// user: The user record to add.
// Returns an error if the operation fails.
func (uc AuthUseCase) Register(ctx context.Context, user entities.User) error {

	if err := uc.Validate(user); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			return formatValidationError(validationErrors)
		}
		return err
	}

	exists, err := uc.repo.CheckUserExists(ctx, user.Email, user.Username)
	if err != nil && !errors.Is(err, errors.New("record not found")) {
		return err
	}
	if exists {
		return errors.New("an account with the given email or username already exists")
	}

	user.ID, err = uc.GenerateUUID()
	if err != nil {
		return err
	}
	user.Password, err = uc.HashPassword(user.Password)
	if err != nil {
		return err
	}

	if err := uc.repo.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

// Login checks the user credentials and logs in the user.
// ctx: The context for the operation.
// userLogin: The user login record to check.
// Returns a string and an error if the operation fails.
func (uc AuthUseCase) Login(ctx context.Context, userLogin entities.UserLogin) (string, error) {

	existingUser, err := uc.repo.ReadByEmail(ctx, userLogin.Email)
	if err != nil {
		return "", err
	}

	if err := uc.ComparePasswords(existingUser.Password, userLogin.Password); err != nil {
		return "", err
	}

	existingUser.Token, err = uc.GenerateBearerToken()
	if err != nil {
		return "", err
	}
	existingUser.Metadata.LastLoginAt = time.Now()

	if err := uc.repo.Update(ctx, existingUser); err != nil {
		return "", err
	}

	return existingUser.Token, nil
}

// formatValidationError formats the validation errors.
// errs: The validation errors to format.
// Returns an error with the formatted validation errors.
func formatValidationError(errs validator.ValidationErrors) error {
	for _, e := range errs {
		switch e.Tag() {
		case "required":
			return fmt.Errorf("%s is required", e.Field())
		case "email":
			return fmt.Errorf("%s must be a valid email address", e.Field())
		case "min":
			return fmt.Errorf("%s must be at least %s characters long", e.Field(), e.Param())
		case "max":
			return fmt.Errorf("%s must be no more than %s characters long", e.Field(), e.Param())
			// Add other cases as needed.
		}
	}
	return fmt.Errorf("validation failed")
}

// Package entities provides the functionality to interact with the user entities of the application.
package entities

import (
	"github.com/google/uuid" // UUID package provides the functionality to generate and use UUIDs.
)

// User struct represents a user entity with fields for the user's ID, username, password, email, metadata, and token.
// ID: The UUID of the user.
// Username: The username of the user. It is unique and required, and must be alphanumeric and between 3 and 20 characters long.
// Password: The password of the user. It is required and must be at least 8 characters long.
// Email: The email of the user. It is unique and required, and must be a valid email address.
// Metadata: The metadata of the user.
// Token: The token of the user. It is optional.
type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default"`
	Username string    `json:"username" gorm:"unique;not null" validate:"required,alphanum,min=3,max=20"`
	Password string    `json:"password" gorm:"size:255" validate:"required,min=8"`
	Email    string    `json:"email" gorm:"unique;not null" validate:"required,email"`
	Metadata Metadata  `json:"metadata" gorm:"embedded;embedded_prefix:meta_"`
	Token    string    `json:"token" gorm:"token" validate:"omitempty"`
}

// UserLogin struct represents a user login entity with fields for the user's email and password.
// Email: The email of the user. It is required and must be a valid email address.
// Password: The password of the user. It is required and must be at least 6 characters long.
type UserLogin struct {
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required,gte=6"`
}

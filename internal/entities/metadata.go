// Package entities provides the functionality to interact with the metadata entities of the application.
package entities

import "time" // Time package provides the functionality to work with time.

// Metadata struct represents a metadata entity with fields for the creation time, update time, and last login time of a user.
// CreatedAt: The creation time of the user. It is automatically set when the user is created.
// UpdatedAt: The update time of the user. It is automatically updated when the user is updated.
// LastLoginAt: The last login time of the user. It is set to null by default.
type Metadata struct {
	CreatedAt   time.Time `json:"created_at" db:"created_at" gorm:"autoCreateTime"`     // The creation time of the user.
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at" gorm:"autoUpdateTime"`     // The update time of the user.
	LastLoginAt time.Time `json:"last_login_at" db:"last_login_at" gorm:"default:null"` // The last login time of the user.
}

// Package auth provides the functionality to interact with user authentication data.
package auth

import "github.com/labstack/echo/v4"

// Handlers is an interface that defines the methods required for handling user authentication operations.
// It includes methods for registering, getting all users, and logging in.
type Handlers interface {
	// Register handles the registration of a new user.
	// Returns an echo.HandlerFunc that handles the HTTP request for user registration.
	Register() echo.HandlerFunc

	// GetAll handles the retrieval of all user records.
	// Returns an echo.HandlerFunc that handles the HTTP request for retrieving all users.
	GetAll() echo.HandlerFunc

	// Login handles the login of a user.
	// Returns an echo.HandlerFunc that handles the HTTP request for user login.
	Login() echo.HandlerFunc
}

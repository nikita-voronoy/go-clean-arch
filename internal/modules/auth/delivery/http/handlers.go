// Package http provides the functionality to handle HTTP requests for the auth module.
package http

import (
	"fmt"
	"github.com/labstack/echo/v4" // Echo is a high performance, extensible, minimalist web framework for Go.
	"net/http"
	"time"
	"x-ci-cd/config"                // Config package provides the functionality to interact with the configuration of the application.
	"x-ci-cd/internal/entities"     // Entities package provides the functionality to interact with the entities of the application.
	"x-ci-cd/internal/modules/auth" // Auth package provides the functionality to interact with the auth module.
)

// AuthHandlers struct represents auth handlers that provide methods for handling HTTP requests for the auth module.
type AuthHandlers struct {
	cfg    *config.Config // The configuration for the auth handlers.
	authUC auth.UseCase   // The auth use case for the auth handlers.
}

// NewAuthHandlers creates new auth handlers with the provided configuration and auth use case.
// cfg: The configuration for the auth handlers.
// authUC: The auth use case for the auth handlers.
// Returns an AuthHandlers object.
func NewAuthHandlers(cfg *config.Config, authUC auth.UseCase) *AuthHandlers {
	return &AuthHandlers{
		cfg:    cfg,
		authUC: authUC,
	}
}

// Register registers a new user.
// @route POST /auth/register
// @group Authentication
// @param {User.model} user.body.required - User details
// @returns {object} 201 - An account has been successfully created.
// @returns {object} 400 - The request could not be understood or was missing required parameters.
// @returns {object} 409 - An account with the given email or username already exists.
func (h *AuthHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user entities.User
		if err := c.Bind(&user); err != nil {
			// Return an HTTP error with status code 400 for bad requests.
			return echo.NewHTTPError(http.StatusBadRequest, "failed to bind user")
		}

		if err := h.authUC.Register(c.Request().Context(), user); err != nil {
			return echo.NewHTTPError(http.StatusConflict, fmt.Sprintf("failed to register user: %v", err))
		}

		return c.JSON(http.StatusCreated, user)
	}
}

// GetAll retrieves all user records.
// @route GET /auth/all
// @group Authentication
// @returns {Array} 200 - An array of user info
// @returns {object} 500 - Server error
func (h *AuthHandlers) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []entities.User
		users, err := h.authUC.GetAll(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to get all users: %v", err))
		}
		return c.JSON(http.StatusOK, users)
	}
}

// Login logs in a user.
// @route POST /auth/login
// @group Authentication
// @param {UserLogin.model} userLogin.body.required - User login details
// @returns {object} 200 - Successful login
// @returns {object} 400 - Invalid username or password
// @returns {object} 401 - Unauthorized access
func (h *AuthHandlers) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var login entities.UserLogin
		if err := c.Bind(&login); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "failed to bind user")
		}

		token, err := h.authUC.Login(c.Request().Context(), login)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("failed to login user: %v", err))
		}

		c.SetCookie(&http.Cookie{
			Name:    "token",
			Value:   token,
			Expires: time.Now().Add(1 * time.Minute),
		})

		return c.JSON(http.StatusOK, fmt.Sprintf("Bearer: %s", token))
	}
}

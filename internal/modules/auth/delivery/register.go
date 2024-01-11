// Package delivery provides the functionality to deliver the responses of the auth module.
package delivery

import (
	"github.com/labstack/echo/v4"                 // Echo is a high performance, extensible, minimalist web framework for Go.
	"x-ci-cd/config"                              // Config package provides the functionality to interact with the configuration of the application.
	"x-ci-cd/internal/modules/auth"               // Auth package provides the functionality to interact with the auth module.
	"x-ci-cd/internal/modules/auth/delivery/http" // HTTP package provides the functionality to deliver the responses of the auth module over HTTP.
)

// AuthDelivery struct represents an auth delivery that provides methods for delivering the responses of the auth module.
// It includes an AuthHandlers object for handling the responses and a function for setting up the routes.
type AuthDelivery struct {
	Handlers        *http.AuthHandlers    // The handlers for the auth responses.
	SetupRoutesFunc func(echo *echo.Echo) // The function for setting up the routes.
}

// NewAuthDelivery creates a new auth delivery with the provided configuration and auth use case.
// cfg: The configuration for the auth delivery.
// uc: The auth use case for the auth delivery.
// Returns an AuthDelivery object.
func NewAuthDelivery(cfg *config.Config, uc auth.UseCase) *AuthDelivery {
	handlers := http.NewAuthHandlers(cfg, uc) // Creates new auth handlers with the provided configuration and auth use case.

	// Returns a new AuthDelivery object with the created handlers and a function for setting up the routes.
	return &AuthDelivery{
		Handlers: handlers,
		SetupRoutesFunc: func(e *echo.Echo) {
			http.MapAuthRoutes(e.Group("/auth"), handlers) // Maps the auth routes to the "/auth" group of the Echo instance.
		},
	}
}

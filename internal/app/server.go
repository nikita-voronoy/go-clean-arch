// Package app provides the functionality to create and manage the server of the application.
package app

import (
	"context"                                // Context package provides the functionality to pass deadlines, cancel signals, and other request-scoped values across API boundaries and between processes.
	"github.com/labstack/echo/v4"            // Echo is a high performance, extensible, minimalist web framework for Go.
	"github.com/labstack/echo/v4/middleware" // Middleware package provides the functionality to use middleware with Echo.
	"go.uber.org/fx"                         // Fx is a framework for Go that provides the tools needed to build a dependency graph and invoke components in the correct order.
	"log"                                    // Log package provides the functionality to implement logging.
	"strconv"                                // Strconv package provides the functionality to convert strings to basic data types.
	"x-ci-cd/config"                         // Config package provides the functionality to interact with the configuration of the application.
)

// NewServer creates a new Echo server with the provided lifecycle and configuration.
// lc: The lifecycle for the server.
// cfg: The configuration for the server.
// The server uses the Logger and Recover middleware from Echo.
// The server starts when the lifecycle starts and shuts down when the lifecycle stops.
// Returns an Echo object.
func NewServer(lc fx.Lifecycle, cfg *config.Config) *echo.Echo {
	// Creates a new Echo instance.
	server := echo.New()

	// Adds the Logger middleware to the Echo instance.
	server.Use(middleware.Logger())

	// Adds the Recover middleware to the Echo instance.
	server.Use(middleware.Recover())

	// Appends a Hook to the lifecycle with OnStart and OnStop functions.
	lc.Append(fx.Hook{
		// The OnStart function starts the Echo server in a new goroutine.
		OnStart: func(ctx context.Context) error {
			go func() {
				// Starts the Echo server with the host and port from the configuration.
				// If the server fails to start, it logs the error.
				if err := server.Start(cfg.Server.Host + ":" + strconv.Itoa(cfg.Server.Port)); err != nil {
					log.Printf("Failed to start Echo app: %v\n", err)
				}
			}()
			return nil
		},
		// The OnStop function shuts down the Echo server.
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	// Returns the Echo instance.
	return server
}

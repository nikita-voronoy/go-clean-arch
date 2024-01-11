// Package main provides the entry point for the application.
package main

import (
	"go.uber.org/fx"                       // Fx is a framework for Go that provides the tools needed to build a dependency graph and invoke components in the correct order.
	"x-ci-cd/config"                       // Config package provides the functionality to interact with the configuration of the application.
	"x-ci-cd/internal/app"                 // App package provides the functionality to create and manage the server of the application.
	"x-ci-cd/internal/modules/auth/module" // Module package provides the functionality to interact with the auth module of the application.
	"x-ci-cd/pkg/database"                 // Database package provides the functionality to interact with the database of the application.
)

// main function is the entry point for the application.
// It creates a new Fx application with the provided providers and modules.
// The providers are the configuration, database, and server of the application.
// The modules are the auth module of the application.
// The application is run with the Run method of Fx.
func main() {
	fx.New(
		fx.Provide(
			config.NewConfig,     // Provides the configuration of the application.
			database.NewDatabase, // Provides the database of the application.
			app.NewServer,        // Provides the server of the application.
		),
		module.Module, // Provides the auth module of the application.
	).Run() // Runs the Fx application.
}

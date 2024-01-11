// Package module provides the functionality to interact with the auth module.
package module

import (
	"github.com/labstack/echo/v4"                 // Echo is a high performance, extensible, minimalist web framework for Go.
	"go.uber.org/fx"                              // Fx is a framework for Go that provides the building blocks for your service architectures.
	"x-ci-cd/internal/modules/auth/delivery"      // Delivery package provides the functionality to deliver the responses of the auth module.
	"x-ci-cd/internal/modules/auth/delivery/http" // HTTP package provides the functionality to deliver the responses of the auth module over HTTP.
	"x-ci-cd/internal/modules/auth/usecase"       // Usecase package provides the functionality to interact with the use cases of the auth module.
	"x-ci-cd/internal/storage/user"               // User package provides the functionality to interact with the user storage.
)

// Module is a Fx options group that provides and invokes the necessary dependencies for the auth module.
var Module = fx.Options(
	fx.Provide(
		user.NewUserRepository,   // Provides a new user repository.
		usecase.NewAuthUC,        // Provides a new auth use case.
		http.NewAuthHandlers,     // Provides new auth handlers.
		delivery.NewAuthDelivery, // Provides a new auth delivery.
	),
	fx.Invoke(registerAuthRoutes), // Invokes the function to register the auth routes.
)

// registerAuthRoutes registers the auth routes with the provided Echo instance and auth handlers.
// e: The Echo instance to register the routes with.
// handlers: The auth handlers to use for the routes.
func registerAuthRoutes(e *echo.Echo, handlers *http.AuthHandlers) {
	http.MapAuthRoutes(e.Group("/auth"), handlers) // Maps the auth routes to the "/auth" group of the Echo instance.
}

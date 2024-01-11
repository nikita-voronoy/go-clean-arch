// Package http provides the functionality to map the routes of the auth module over HTTP.
package http

import (
	"github.com/labstack/echo/v4"   // Echo is a high performance, extensible, minimalist web framework for Go.
	"x-ci-cd/internal/modules/auth" // Auth package provides the functionality to interact with the auth module.
)

// MapAuthRoutes maps the auth routes to the provided Echo group with the provided auth handlers.
// authGroup: The Echo group to map the routes to.
// h: The auth handlers to use for the routes.
// The routes include:
// POST /register: Registers a new user. Expects a JSON body with the user details.
// POST /login: Logs in a user. Expects a JSON body with the user login details.
// GET /all: Retrieves all user records.
func MapAuthRoutes(authGroup *echo.Group, h auth.Handlers) {
	// @route POST /auth/register
	// @group Authentication
	// @param {User.model} user.body.required - User details
	// @returns {object} 200 - An account has been successfully created.
	// @returns {object} 400 - The request could not be understood or was missing required parameters.
	// @returns {object} 500 - Server error
	authGroup.POST("/register", h.Register())

	// @route POST /auth/login
	// @group Authentication
	// @param {UserLogin.model} userLogin.body.required - User login details
	// @returns {object} 200 - Successful login
	// @returns {object} 400 - Invalid username or password
	// @returns {object} 500 - Server error
	authGroup.POST("/login", h.Login())

	// @route GET /auth/all
	// @group Authentication
	// @returns {Array} 200 - An array of user info
	// @returns {object} 500 - Server error
	authGroup.GET("/all", h.GetAll())
}

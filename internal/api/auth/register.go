package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisterResponse struct {
	Username   string `json:"username"`
	Registered bool   `json:"registered"`
}

// Register a user, given a username, password, and email.
func RegisterHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}

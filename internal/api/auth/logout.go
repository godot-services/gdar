package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LogoutRequest struct {
	Token string `json:"token"`
}

type LogoutResponse struct {
	Authenticated bool   `json:"authenticated"`
	Token         string `json:"token"`
}

// Logout a user, given a token. The token is invalidated in the process.
func LogoutHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}

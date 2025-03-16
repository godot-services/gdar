package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Authenticated bool   `json:"authenticated"`
	Username      string `json:"username"`
	Token         string `json:"token"`
}

// Login as a given user. Results in a token which can be used for authenticated requests.
func LoginHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}

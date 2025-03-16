package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ChangePasswordRequest struct {
	Token       string `json:"token"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangePasswordResponse struct {
	Token string `json:"token"`
}

// Change a user's password. The token is invalidated in the process.
func ChangePasswordHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}

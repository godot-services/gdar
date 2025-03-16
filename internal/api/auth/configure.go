package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ConfigureResponse struct {
	Categories []Category `json:"categories"`
	Token      string     `json:"token"`
	LoginURL   string     `json:"login_url"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// Get a list of categories (needed for filtering assets) and potentially a login URL
// which can be given to the user in order to authenticate him in the engine
// (currently unused and not working).
func ConfigureHandler(c echo.Context) error {
	// TODO implement this
	return c.NoContent(http.StatusNotImplemented)
}

package main

import (
	"net/http"

	"github.com/godot-services/gdar/internal/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Static("/", "public")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// TODO asset frontend
	// TODO management frontend
	// TODO health api

	apiGroup := e.Group("/api")
	api.Routes(apiGroup)

	e.Logger.Fatal(e.Start(":1323"))
}

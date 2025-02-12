package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) error {
	e.GET("/", MainPage)
	e.GET("/users", UsersPage)
	return nil
}

func MainPage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func UsersPage(c echo.Context) error {
	return c.String(http.StatusOK, "Users")
}

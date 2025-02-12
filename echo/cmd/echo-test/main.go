package main

import (
	"echo-test/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	handler.Register(e)

	e.Logger.Fatal(e.Start(":1323"))
}

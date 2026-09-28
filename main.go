package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func apply(c *echo.Context) error {
	handle := c.FormValue("handle")
	email := c.FormValue("email")
	bio := c.FormValue("bio")

	fmt.Printf("Handle: %s\n Email: %s\n Bio: %s", handle, email, bio)
	return c.Redirect(http.StatusFound, "/")
}

func main() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.Static("/assets", "static")

	e.GET("/", func(c *echo.Context) error {
		return c.File("html/home.html")
	})
	e.GET("/faq", func(c *echo.Context) error {
		return c.File("html/faq.html")
	})
	e.GET("/privacy", func(c *echo.Context) error {
		return c.File("html/privacy.html")
	})

	e.POST("/apply", apply)

	if err := e.Start(":42069"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

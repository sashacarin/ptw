package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func apply(c *echo.Context) error {
	handle := c.FormValue("handle")
	email := c.FormValue("email")
	bio := c.FormValue("message")
	err := sendEmail(handle, email, bio)
	if err != nil {
		log.Printf("Email send fail: %v", err)
		c.Redirect(http.StatusFound, "/apply-error")
	}
	fmt.Printf("Handle: %s\n Email: %s\n Bio: %s", handle, email, bio)
	return c.Redirect(http.StatusFound, "/apply-success")
}

func sendEmail(handle string, email string, bio string) error {
	from := os.Getenv("PTW_SENDER")
	receiver := os.Getenv("PTW_RECEIVER")
	password := os.Getenv("PTW_PASSWORD")

	addr := "smtp.migadu.com:587"
	host := "smtp.migadu.com"

	to := []string{receiver}

	msg := []byte(
		"From: " + from + "\r\n" +
			"To: " + receiver + "\r\n" +
			"Subject: PTW Application\r\n" +
			"\r\n" +
			"Handle: " + handle + "\r\n" +
			"Email: " + email + "\r\n" +
			"Bio: " + bio + "\r\n",
	)

	auth := smtp.PlainAuth("", from, password, host)

	if err := smtp.SendMail(addr, auth, from, to, msg); err != nil {
		return fmt.Errorf("send email: %w", err)
	}

	fmt.Println("Email sent successfully")
	return nil
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
	e.GET("/apply-error", func(c *echo.Context) error {
		return c.File("html/apply_error.html")
	})
	e.GET("/apply-success", func(c *echo.Context) error {
		return c.File("html/apply_success.html")
	})

	e.POST("/apply", apply)

	if err := e.Start(":42069"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

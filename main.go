package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/user/:name", func(c echo.Context) error {
		return c.String(http.StatusOK,fmt.Sprintf("Hello, %s!",c.Param("name")))
	})
	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK,[]string{"Alice","Bob"})
	})
	e.Logger.Fatal(e.Start(":8080"))
}

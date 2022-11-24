package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func main() {

	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "!Pong!")
	})

	e.GET("/user/:id", getUser)

	e.Logger.Fatal(e.Start(":9999"))

}

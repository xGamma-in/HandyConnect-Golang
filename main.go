package main

import (
	"HandyConnectBackend/api/v1/users"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})
	e.GET("/users/:id", users.GetUser)
	e.Logger.Fatal(e.Start(":8080"))
}

package main

import (
	"net/http"
	"time"

	myapi "./api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/get-hello1", getTimeNow)
	e.GET("/get-hello2", gagaulala)

	go func() {
		time.Sleep(10 * time.Second)

		myapi.SetAPIRequest()

	}()

	e.Start(":8080")

}

func getTimeNow(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

func gagaulala(c echo.Context) error {
	return c.String(http.StatusOK, "gagaulala")
}

package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"bitbucket.org/kyicy/seifer/config"
)

func main() {
	config.Load()

	conf := config.Get()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf("%v:%v", conf.Addr, conf.Port)))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, my World !!!")
}

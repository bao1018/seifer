package main

import (
	"fmt"

	"bitbucket.org/kyicy/seifer/app/model"

	"bitbucket.org/kyicy/seifer/app"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"bitbucket.org/kyicy/seifer/config"
)

func main() {
	config.Load()

	conf := config.Get()

	db, err := gorm.Open("postgres", conf.Database.Connection)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	model.Init(db)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	app.RegisterRoute(e)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf("%v:%v", conf.Addr, conf.Port)))
}

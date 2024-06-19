package main

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"gohome/controllers"
	"log"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := createApp()
	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func createApp() *pocketbase.PocketBase {
	app := pocketbase.New()
	app.OnBeforeServe().Add(setupWebsite)
	return app
}

func setupWebsite(e *core.ServeEvent) error {
	e.Router.Static("/assets", "assets")

	e.Router.GET("/",
    func(c echo.Context) error {
		return controllers.HomeController(e.App, c)
	})

	e.Router.GET("/news", 
    func(c echo.Context) error {
		return controllers.NewsListController(e.App, c)
	})

	e.Router.GET("/news/:id", 
    func(c echo.Context) error {
		return controllers.NewsItemController(e.App, c)
	})

	e.Router.GET("/me", 
    func(c echo.Context) error {
		return controllers.MeController(e.App, c)
	})
	return nil
}

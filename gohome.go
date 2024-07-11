package main

import (
	"embed"
	"gohome/controllers"
	"log"
	"net/http"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

//go:embed  assets/*
var assets embed.FS

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
	setupStatic(e)

	e.Router.GET("/",
		func(c echo.Context) error {
			return controllers.HomeController(e.App, c)
		})

	e.Router.GET("/projects",
		func(c echo.Context) error {
			return controllers.ProjectsController(e.App, c)
		})

	e.Router.GET("/feed",
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

func setupStatic(e *core.ServeEvent) error {
	var contentHandler = echo.WrapHandler(http.FileServer(http.FS(assets)))
	e.Router.GET("/assets/*", contentHandler)

	return nil
}

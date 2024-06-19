package main

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"

	// "github.com/pocketbase/pocketbase/apis"
	"gohome/controllers"
	"gohome/templates"
	"log"

	"github.com/pocketbase/pocketbase/core"
	// "net/http"
	// "os"
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
	component := templates.Base()
	e.Router.Static("/assets", "assets")
	e.Router.GET("/", echo.WrapHandler(templ.Handler(component)))

	e.Router.GET("/news/:name", 
    func(c echo.Context) error {
		return controllers.NewsItemController(e.App, c)
	})
	// e.Router.GET("/news/:id", NewsItemController echo.WrapHandler(templ.Handler(controllers.NewsItemController(&e.App))))
	// e.Router.GET("/hello", func(c echo.Context) error { return c.String(http.StatusOK, "Hello World") })
	return nil
}

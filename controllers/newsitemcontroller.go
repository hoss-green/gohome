package controllers

import (
	"gohome/data/models"
	"gohome/templates"
	"log"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

// func(c echo.Context) error
func NewsItemController(e core.App, c echo.Context) error {
	name := c.PathParam("name")
	// templ.Handler(controllers.NewsItemController(&e.App)
	// pbApp := a
	// echo.WrapHandler(templ.Handler(controllers.NewsItemController(&e.App)

	newsitem := models.NewsItem{}

	err := e.Dao().DB().NewQuery("SELECT id, created, updated, title, content FROM news_items").One(&newsitem)

	if err != nil {
		log.Printf("Select Error: %s\r\n", err)
		val, err := templ.ToGoHTML(c.Request().Context(), templates.Error())
		if err != nil {
			log.Printf("Templ Error: %s\r\n", err)
		}
		return c.String(500, string(val))
	}

	val, err := templ.ToGoHTML(c.Request().Context(), templates.NewsItem(name))
	if err != nil {

	}
	// (templates.Error())
	// return c.String(404, string(val))
	return c.HTML(200, string(val))
	// c.String()
}

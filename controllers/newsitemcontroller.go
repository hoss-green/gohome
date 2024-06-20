package controllers

import (
	"gohome/models"
	"gohome/templates"
	"log"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func NewsItemController(e core.App, c echo.Context) error {
	id := c.PathParam("id")

	newsitem := models.NewsItem{}

	err := e.Dao().DB().NewQuery("SELECT id, created, updated, title, content FROM news_items WHERE id = {:id}").
		Bind(dbx.Params{
			"id": id,
		}).
		One(&newsitem)

	if err != nil {
		log.Printf("Select Error: %s\r\n", err)
		val, err := templ.ToGoHTML(c.Request().Context(), templates.Error())
		if err != nil {
			log.Printf("Templ Error: %s\r\n", err)
		}
		return c.String(500, string(val))
	}

	return render(c, templates.NewsItem(newsitem), false)
	// c.String()
}

package controllers

import (
	"gohome/models"
	"gohome/templates"
	"log"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func NewsListController(e core.App, c echo.Context) error {
	newsListItems := []models.NewsListItem{}

	err := e.Dao().DB().NewQuery("SELECT id, title FROM news_items").All(&newsListItems)

	if err != nil {
		log.Printf("Select Error: %s\r\n", err)
		val, err := templ.ToGoHTML(c.Request().Context(), templates.Error())
		if err != nil {
			log.Printf("Templ Error: %s\r\n", err)
		}
		return c.String(500, string(val))
	}

	return render(c, templates.NewsList(newsListItems))
}

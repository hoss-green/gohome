package controllers

import (
	"gohome/models"
	"gohome/templates"
	"log"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func MeController(e core.App, c echo.Context) error {
	pageItem := models.PageItem{}

	err := e.Dao().DB().NewQuery("SELECT id, created, updated, title, content FROM pages WHERE page_name='me'").
		One(&pageItem)

	if err != nil {
		log.Printf("Select Error: %s\r\n", err)
		val, err := templ.ToGoHTML(c.Request().Context(), templates.Error())
		if err != nil {
			log.Printf("Templ Error: %s\r\n", err)
		}
		return c.String(500, string(val))
	}
	return render(c, templates.MePage(pageItem), false)
}

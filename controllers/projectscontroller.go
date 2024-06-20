package controllers

import (
	"gohome/models"
	"gohome/templates"
	"log"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func ProjectsController(e core.App, c echo.Context) error {
	projectItems := []models.ProjectItem{}

	err := e.Dao().DB().NewQuery("SELECT id, title, summary, link FROM projects").All(&projectItems)

	if err != nil {
		log.Printf("Select Error: %s\r\n", err)
		val, err := templ.ToGoHTML(c.Request().Context(), templates.Error())
		if err != nil {
			log.Printf("Templ Error: %s\r\n", err)
		}
		return c.String(500, string(val))
	}

	return render(c, templates.Projects(projectItems), false)
}

package controllers

import (
	"gohome/models"
	"gohome/templates"
	"log"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func HomeController(e core.App, c echo.Context) error {
	projectItems := []models.ProjectItem{}
	newsItem := models.NewsItem{}

	err1 := e.Dao().DB().NewQuery("SELECT id, created, title, summary, link FROM projects ORDER BY created DESC LIMIT 3").All(&projectItems)
	err2 := e.Dao().DB().NewQuery("SELECT id, created, updated, summary, title, content FROM news_items ORDER BY created DESC LIMIT 1").
		One(&newsItem)
	if err1 != nil || err2 != nil {
		log.Printf("err1 Error: %s\r\n", err1)
		log.Printf("err2 Error: %s\r\n", err2)
		val, err := templ.ToGoHTML(c.Request().Context(), templates.Error())
		if err != nil {
			log.Printf("Templ Error: %s\r\n", err)
		}
		return c.String(500, string(val))
	}

  // if len(newsItem.Content) > 400 {
  //   newsItem.Content = fmt.Sprintf("%s...", newsItem.Content[:350])
  // }
	// return render(c, templates.Projects(projectItems), false)
	return render(c, templates.HomePage(projectItems, newsItem), true)
}

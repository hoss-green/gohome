package controllers

import (
	// "gohome/models"
	"gohome/templates"
	// "log"

	// "github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func HomeController(e core.App, c echo.Context) error {
	return render(c, templates.HomePage(), true)
}

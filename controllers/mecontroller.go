package controllers

import (
	"gohome/templates"

	// "github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

func MeController(_ core.App, c echo.Context) error {
	return render(c, templates.HomePage())
}

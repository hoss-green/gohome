package controllers

import (
	// _ "embed"
	"gohome/templates"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
)


func render(c echo.Context, bodyComponent templ.Component, fullwidth bool) error {
  layout, err := templ.ToGoHTML(c.Request().Context(), templates.Layout(bodyComponent, fullwidth))
	if err != nil {
		val, _ := templ.ToGoHTML(c.Request().Context(), templates.Error())
		return c.String(500, string(val))
	}
	return c.HTML(200, string(layout))
}

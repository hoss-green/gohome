package controllers

import (
	"gohome/templates"
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func ContactController(e core.App, c echo.Context) error {
	qp := c.QueryParam("success")
	if qp == "ok" {
		return render(c, templates.Contact("Message Sent Succesfully"), true)
	}
	return render(c, templates.Contact(""), true)
}

func SendContactControllerMessage(e core.App, c echo.Context) error {

	name := c.FormValue("name")
	message := c.FormValue("message")
	email := c.FormValue("email")
	summary := c.FormValue("summary")

	_, err := e.Dao().DB().NewQuery("INSERT into messages (name, message, email, summary) VALUES ({:name}, {:message}, {:email}, {:summary})").
		Bind(dbx.Params{
			"name":    name,
			"message": message,
			"email":   email,
			"summary": summary,
		}).Execute()
	if err != nil {
		log.Println(err.Error())
		return render(c, templates.Contact("message failed to send, please try again"), true)
	}

	return c.Redirect(http.StatusSeeOther, "/contact?success=ok")
	// return render(c, templates.Contact("message sent"), true)
}

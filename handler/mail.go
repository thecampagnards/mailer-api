package handler

import (
	"mailer-api/config"
	"mailer-api/dao"
	"mailer-api/types"

	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// Mail struct which contains the functions of this class
type Mail struct {
}

// Send rest function which send a mail
func (m *Mail) Send(c echo.Context) error {

	var mail types.Mail
	var err error

	// Retreiving the smtp conf
	mail.SMTP, err = dao.GetByIDSMTPServer(c.QueryParam("server-smtp-id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Retreiving the template conf
	mail.Template, err = dao.GetByIDMailTemplate(c.QueryParam("template-id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// formating the readcloser to stringas template vars
	b := new(bytes.Buffer)
	b.ReadFrom(c.Request().Body)
	err = json.Unmarshal(b.Bytes(), &mail.TemplateVars)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	splitFn := func(c rune) bool {
		return c == ','
	}

	// Use FieldsFunc because Split create empty slice
	mail.To = strings.FieldsFunc(c.QueryParam("to"), splitFn)
	mail.Cc = strings.FieldsFunc(c.QueryParam("cc"), splitFn)
	mail.Cci = strings.FieldsFunc(c.QueryParam("cci"), splitFn)
	/*
		mail.Attachement, err := c.FormFile("attachement")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	*/

	mailer := config.Mailer{}
	// Send the mail
	err = mailer.Send(mail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "send")
}

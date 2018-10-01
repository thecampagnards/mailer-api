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

	// Check if multipart
	if strings.HasPrefix(c.Request().Header.Get(echo.HeaderContentType), echo.MIMEMultipartForm) {

		// get form data
		form, err := c.MultipartForm()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// formating the string vars to object template vars
		if json.Unmarshal([]byte(form.Value[types.FORM_DATA_DATA_FIELD_NAME][0]), &mail.TemplateVars) != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		cf := config.File{}

		for _, file := range form.File[types.FORM_DATA_ATTACHMENTS_FIELD_NAME] {
			// Save the file
			fn, err := cf.Save(file)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			// defer the remove of the file (will be execute after a return)
			defer cf.Remove(fn)

			// Add the filename in the attachments
			mail.Attachments = append(mail.Attachments, fn)
		}

	} else {
		// formating the readcloser to strings template vars
		b := new(bytes.Buffer)
		b.ReadFrom(c.Request().Body)
		if json.Unmarshal(b.Bytes(), &mail.TemplateVars) != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
	}

	splitFn := func(c rune) bool {
		return c == ','
	}

	// Use FieldsFunc because Split create empty slice
	mail.To = strings.FieldsFunc(c.QueryParam("to"), splitFn)
	mail.Cc = strings.FieldsFunc(c.QueryParam("cc"), splitFn)
	mail.Cci = strings.FieldsFunc(c.QueryParam("cci"), splitFn)

	mailer := config.Mailer{}
	// Send the mail
	err = mailer.Send(mail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "send")
}

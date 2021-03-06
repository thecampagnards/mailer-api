package handler

import (
	"mailer-api/config"
	"mailer-api/dao"
	"mailer-api/types"

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
		c.Logger().Errorf("Error when retreiving smtp server : %s", c.QueryParam("server-smtp-id"))
		return c.JSON(http.StatusBadRequest, err)
	}

	// Retreiving the template conf
	mail.Template, err = dao.GetByIDMailTemplate(c.QueryParam("template-id"))
	if err != nil {
		c.Logger().Errorf("Error when retreiving template : %s", c.QueryParam("template-id"))
		return c.JSON(http.StatusBadRequest, err)
	}

	if mail.Template.TemplateURL != "" {
		mail.Template.Template, err = config.FetchBodyFromURL(mail.Template.TemplateURL)
		if err != nil {
			c.Logger().Errorf("Error when requesting the url : %s", mail.Template.TemplateURL)
			return c.JSON(http.StatusBadRequest, err)
		}
	}

	for _, ID := range mail.Template.LayoutIDs {

		// Retreiving the layout
		layout, err := dao.GetByIDLayout(ID.Hex())
		if err != nil {
			c.Logger().Errorf("Error when retreiving layout : %s", ID)
			return c.JSON(http.StatusBadRequest, err)
		}

		if layout.LayoutURL != "" {
			layout.Layout, err = config.FetchBodyFromURL(layout.LayoutURL)
			if err != nil {
				c.Logger().Errorf("Error when requesting the url : %s", mail.Template.TemplateURL)
				return c.JSON(http.StatusBadRequest, err)
			}
		}

		mail.Template.Template = layout.Layout + mail.Template.Template
	}

	c.Logger().Info("Template : %s", mail.Template.Template)
	c.Logger().Info("POST request as %s", c.Request().Header.Get(echo.HeaderContentType))

	// Check if multipart
	if strings.HasPrefix(c.Request().Header.Get(echo.HeaderContentType), echo.MIMEMultipartForm) {

		// get form data
		form, err := c.MultipartForm()
		if err != nil {
			c.Logger().Error("Error when parsing the form")
			return c.JSON(http.StatusBadRequest, err)
		}

		// formating the string vars to object template vars
		if json.Unmarshal([]byte(form.Value[types.FORM_DATA_DATA_FIELD_NAME][0]), &mail.TemplateVars) != nil {
			c.Logger().Errorf("Error when parsing to json the field %s", types.FORM_DATA_DATA_FIELD_NAME)
			return c.JSON(http.StatusBadRequest, err)
		}

		cf := config.File{}

		for _, file := range form.File[types.FORM_DATA_ATTACHMENTS_FIELD_NAME] {
			// Save the file
			fn, err := cf.Save(file)
			if err != nil {
				c.Logger().Errorf("Error when saving the file %s", file.Filename)
				return c.JSON(http.StatusBadRequest, err)
			}
			// defer the remove of the file (will be execute after a return)
			defer cf.Remove(fn)

			// Add the filename in the attachments
			mail.Attachments = append(mail.Attachments, fn)
		}

	} else {
		// formating the readcloser to strings template vars
		if err := config.Convert(c.Request().Body, &mail.TemplateVars); err != nil {
			c.Logger().Error("Error when parsing to json body")
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
		c.Logger().Error("Error when sending mail")
		return c.JSON(http.StatusBadRequest, err)
	}

	c.Logger().Info("Mail sended")
	return c.JSON(http.StatusOK, "send")
}

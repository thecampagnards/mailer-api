package handler

import (
	"mailer-api/config"
	"mailer-api/dao"
	"mailer-api/types"

	"net/http"

	"github.com/labstack/echo"
)

// Template struct which contains the functions of this class
type Template struct {
}

// GetAll find all
func (te *Template) GetAll(c echo.Context) error {
	t, err := dao.AllMailTemplates()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, t)
}

// GetByID find one by id
func (te *Template) GetByID(c echo.Context) error {
	t, err := dao.GetByIDMailTemplate(c.Param(":ID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, t)
}

// Save a mail template
func (te *Template) Save(c echo.Context) error {
	var t types.MailTemplate
	err := config.Convert(c.Request().Body, &t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	s, err := dao.CreateorUpdateMailTemplate(t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, s)
}

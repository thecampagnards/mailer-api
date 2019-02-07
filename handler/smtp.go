package handler

import (
	"mailer-api/config"
	"mailer-api/dao"
	"mailer-api/types"

	"net/http"

	"github.com/labstack/echo"
)

// SMTP struct which contains the functions of this class
type SMTP struct {
}

// GetAll find all
func (st *SMTP) GetAll(c echo.Context) error {
	s, err := dao.AllSMTPServers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, s)
}

// GetByID find one by id
func (st *SMTP) GetByID(c echo.Context) error {
	s, err := dao.GetByIDSMTPServer(c.Param(":ID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, s)
}

// Save a smtp server
func (st *SMTP) Save(c echo.Context) error {
	var u types.SMTPServer
	err := config.Convert(c.Request().Body, &u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	s, err := dao.CreateOrUpdateSMTPServer(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, s)
}

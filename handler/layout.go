package handler

import (
	"mailer-api/config"
	"mailer-api/dao"
	"mailer-api/types"

	"net/http"

	"github.com/labstack/echo"
)

// Layout struct which contains the functions of this class
type Layout struct {
}

// GetAll find all
func (te *Layout) GetAll(c echo.Context) error {
	t, err := dao.AllLayouts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, t)
}

// GetByID find one by id
func (te *Layout) GetByID(c echo.Context) error {
	t, err := dao.GetByIDLayout(c.Param(":ID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, t)
}

// Save a layout
func (te *Layout) Save(c echo.Context) error {
	var t types.Layout
	err := config.Convert(c.Request().Body, &t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	s, err := dao.CreateOrUpdateLayout(t)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, s)
}

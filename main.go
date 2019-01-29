package main

import (
	"mailer-api/handler"

	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == os.Getenv("ADMIN_USER") && password == os.Getenv("ADMIN_PASSWORD") {
			e.Logger.Info("Good credentials")
			return true, nil
		}
		e.Logger.Error("Wrong credentials")
		return false, nil
	}))

	SMTP := handler.SMTP{}
	Template := handler.Template{}
	Mail := handler.Mail{}

	conf := e.Group("/configuration")
	// For smtp server conf
	stmp := conf.Group("/smtp")
	stmp.GET("/:id", SMTP.GetByID)
	stmp.GET("", SMTP.GetAll)
	stmp.POST("", SMTP.Save)

	// For template conf
	template := conf.Group("/template")
	template.GET("/:id", Template.GetByID)
	template.GET("", Template.GetAll)
	template.POST("", Template.Save)

	e.POST("/send", Mail.Send)

	e.Logger.Fatal(e.Start(":8080"))
}

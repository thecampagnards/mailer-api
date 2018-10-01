package config

import (
	"mailer-api/types"

	"bytes"
	"crypto/tls"
	"errors"
	"text/template"

	"gopkg.in/gomail.v2"
)

// Mailer struct which contains the functions of this class
type Mailer struct {
}

// Generate the html mail with go template
func (ma *Mailer) Generate(t string, templateVars interface{}) (string, error) {

	tmpl, err := template.New("template").Parse(t)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, templateVars)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

// Send an email
func (ma *Mailer) Send(mail types.Mail) error {

	body, err := ma.Generate(mail.Template.Template, mail.TemplateVars)
	if err != nil {
		return errors.New("There was an error trying to generate the mail with the template: " + err.Error())
	}

	m := gomail.NewMessage()

	m.SetHeader("From", mail.SMTP.From)
	m.SetHeader("To", mail.To...)
	if len(mail.Cc) > 0 {
		m.SetHeader("Cc", mail.Cc...)
	}
	if len(mail.Cci) > 0 {
		m.SetHeader("Cci", mail.Cci...)
	}
	for _, at := range mail.Attachments {
		m.Attach(at)
	}
	m.SetHeader("Subject", mail.Template.Subject)
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer(mail.SMTP.Host, mail.SMTP.Port, mail.SMTP.User, mail.SMTP.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: mail.SMTP.InsecureSkipVerify}

	return d.DialAndSend(m)
}

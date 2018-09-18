package dao

import (
	"mailer-api/config"
	"mailer-api/types"

	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
)

const colTemplate string = "templates"

// AllMailTemplates get all mail templates
func AllMailTemplates() (types.MailTemplates, error) {
	db := config.DB{}
	t := types.MailTemplates{}

	s, err := db.DoDial()

	if err != nil {
		return t, errors.New("There was an error trying to connect with the DB")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(colTemplate)

	err = c.Find(bson.M{}).All(&t)

	if err != nil {
		return t, errors.New("There was an error trying to find the templates")
	}

	return t, err
}

// GetByIDMailTemplate find one mail template by id
func GetByIDMailTemplate(id string) (types.MailTemplate, error) {
	db := config.DB{}
	t := types.MailTemplate{}

	s, err := db.DoDial()

	if err != nil {
		return t, errors.New("There was an error trying to connect with the DB")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(colTemplate)

	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&t)

	if err != nil {
		return t, errors.New("There was an error trying to find the template")
	}

	return t, err
}

// NewMailTemplate Save a mail template
func NewMailTemplate(t types.MailTemplate) (types.MailTemplate, error) {
	db := config.DB{}
	t.ID = bson.NewObjectId()
	t.CreatedAt = time.Now()

	s, err := db.DoDial()

	if err != nil {
		return t, errors.New("There was an error trying to connect to the DB")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(colTemplate)

	err = c.Insert(t)

	if err != nil {
		return t, errors.New("There was an error trying to insert the template to the DB")
	}

	return t, err
}

// DeleteMailTemplate a template mail by id
func DeleteMailTemplate(id string) error {
	db := config.DB{}

	s, err := db.DoDial()

	if err != nil {
		return errors.New("There was an error trying to connect with the DB")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(colTemplate)

	err = c.RemoveId(bson.ObjectIdHex(id))

	if err != nil {
		return errors.New("There was an error trying to remove the template")
	}

	return err
}

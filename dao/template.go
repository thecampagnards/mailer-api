package dao

import (
	"mailer-api/config"
	"mailer-api/types"

	"errors"

	"github.com/globalsign/mgo"
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

// CreateorUpdateMailTemplate Save a mail template
func CreateorUpdateMailTemplate(t types.MailTemplate) (types.MailTemplate, error) {

	if t.TemplateURL != "" {
		if t.Template != "" {
			return t, errors.New("You can put both template and templateURL")
		}

		if config.IsValidURL(t.TemplateURL) {
			return t, errors.New("Your URL is not valid")
		}
	}

	db := config.DB{}

	s, err := db.DoDial()

	if err != nil {
		return t, errors.New("There was an error trying to connect to the DB")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(colTemplate)

	if !t.ID.Valid() {
		t.ID = bson.NewObjectId()
	}

	change := mgo.Change{
		Update:    bson.M{"$set": t},
		ReturnNew: true,
		Upsert:    true,
	}
	_, err = c.FindId(t.ID).Apply(change, &t)

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

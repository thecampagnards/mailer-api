package dao

import (
	"mailer-api/config"
	"mailer-api/types"

	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/imdario/mergo"
)

const colSMTP string = "smtp"

// AllSMTPServers get all smtp servers
func AllSMTPServers() (types.SMTPServers, error) {
	db := config.DB{}
	t := types.SMTPServers{}

	s, err := db.DoDial()

	if err != nil {
		return t, errors.New("There was an error trying to connect with the DB")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(colSMTP)

	err = c.Find(bson.M{}).All(&t)

	if err != nil {
		return t, errors.New("There was an error trying to find the SMTPServers")
	}

	return t, err
}

// GetByIDSMTPServer get one smtp server by id
func GetByIDSMTPServer(id string) (types.SMTPServer, error) {
	db := config.DB{}
	t := types.SMTPServer{}

	s, err := db.DoDial()

	if err != nil {
		return t, errors.New("There was an error trying to connect with the DB")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(colSMTP)

	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&t)

	if err != nil {
		return t, errors.New("There was an error trying to find the SMTPServer")
	}

	return t, err
}

// CreateOrUpdateSMTPServer create or update smtp conf
func CreateOrUpdateSMTPServer(t types.SMTPServer) (types.SMTPServer, error) {
	db := config.DB{}

	s, err := db.DoDial()

	if err != nil {
		return t, errors.New("There was an error trying to connect to the DB")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(colSMTP)

	err = c.Insert(t)

	if t.ID.Valid() {
		smtp, _ := GetByIDSMTPServer(t.ID.Hex())
		if err := mergo.Merge(&t, smtp); err != nil {
			return t, err
		}
		err = c.UpdateId(t.ID, t)
	} else {
		t.ID = bson.NewObjectId()
		t.CreatedAt = time.Now()
		err = c.Insert(t)
	}

	if err != nil {
		return t, errors.New("There was an error trying to insert the SMTPServer to the DB")
	}

	return t, err
}

// DeleteSMTPServers remove a stmp conf by id
func DeleteSMTPServers(id string) error {
	db := config.DB{}

	s, err := db.DoDial()

	if err != nil {
		return errors.New("There was an error trying to connect with the DB")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(colSMTP)

	err = c.RemoveId(bson.ObjectIdHex(id))

	if err != nil {
		return errors.New("There was an error trying to remove the SMTPServer")
	}

	return err
}

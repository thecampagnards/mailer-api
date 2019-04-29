package dao

import (
	"mailer-api/config"
	"mailer-api/types"

	"errors"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
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

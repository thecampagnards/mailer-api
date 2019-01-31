package dao

import (
	"mailer-api/config"
	"mailer-api/types"

	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
)

const colLayout string = "layouts"

// AllLayouts get all layouts
func AllLayouts() (types.Layouts, error) {
	db := config.DB{}
	t := types.Layouts{}

	s, err := db.DoDial()

	if err != nil {
		return t, errors.New("There was an error trying to connect with the DB")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(colLayout)

	err = c.Find(bson.M{}).All(&t)

	if err != nil {
		return t, errors.New("There was an error trying to find the layouts")
	}

	return t, err
}

// GetByIDLayout find one layout by id
func GetByIDLayout(id string) (types.Layout, error) {
	db := config.DB{}
	t := types.Layout{}

	s, err := db.DoDial()

	if err != nil {
		return t, errors.New("There was an error trying to connect with the DB")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(colLayout)

	err = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&t)

	if err != nil {
		return t, errors.New("There was an error trying to find the layout")
	}

	return t, err
}

// NewLayout Save a layout
func NewLayout(t types.Layout) (types.Layout, error) {
	db := config.DB{}
	t.ID = bson.NewObjectId()
	t.CreatedAt = time.Now()

	s, err := db.DoDial()

	if err != nil {
		return t, errors.New("There was an error trying to connect to the DB")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(colLayout)

	err = c.Insert(t)

	if err != nil {
		return t, errors.New("There was an error trying to insert the layout to the DB")
	}

	return t, err
}

// DeleteLayout a layout by id
func DeleteLayout(id string) error {
	db := config.DB{}

	s, err := db.DoDial()

	if err != nil {
		return errors.New("There was an error trying to connect with the DB")
	}

	defer s.Close()

	c := s.DB(db.Name()).C(colLayout)

	err = c.RemoveId(bson.ObjectIdHex(id))

	if err != nil {
		return errors.New("There was an error trying to remove the layout")
	}

	return err
}

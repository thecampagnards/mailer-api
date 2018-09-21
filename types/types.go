package types

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type MailTemplate struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"createdAt,omitempty" bson:"createdAt"`
	Template  string
	Subject   string
	// The variables below are just for information
	// this is not used in sending mail
	Variables   interface{}
	Description string
}

type MailTemplates []MailTemplate

type SMTPServer struct {
	ID                 bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt          time.Time     `json:"createdAt,omitempty" bson:"createdAt"`
	Host               string
	Port               int
	User               string
	Password           string
	From               string
	InsecureSkipVerify bool
}

type SMTPServers []SMTPServer

type Mail struct {
	To           []string
	Cc           []string
	Cci          []string
	Attachement  string
	TemplateVars interface{}
	Template     MailTemplate
	SMTP         SMTPServer
}

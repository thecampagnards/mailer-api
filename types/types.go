package types

import (
	"github.com/globalsign/mgo/bson"
)

const (
	FORM_DATA_ATTACHMENTS_FIELD_NAME = "attachments"
	FORM_DATA_DATA_FIELD_NAME        = "data"
)

type MailTemplate struct {
	ID          bson.ObjectId   `yaml:"_id,omitempty" json:"_id,omitempty" bson:"_id,omitempty"`
	Template    string          `yaml:"template" json:"template" bson:"template"`
	TemplateURL string          `yaml:"template_url" json:"template_url" bson:"template_url"`
	Subject     string          `yaml:"subject" json:"subject" bson:"subject"`
	LayoutIDs   []bson.ObjectId `yaml:"layout_ids" json:"layout_ids" bson:"layout_ids"`
	// The variables below are just for information
	// this is not used in sending mail
	Variables   interface{} `yaml:"variables" json:"variables" bson:"variables"`
	Description string      `yaml:"description" json:"description" bson:"description"`
}

type MailTemplates []MailTemplate

type Layout struct {
	ID        bson.ObjectId `yaml:"_id,omitempty" json:"_id,omitempty" bson:"_id,omitempty"`
	Layout    string        `yaml:"layout" json:"layout" bson:"layout"`
	LayoutURL string        `yaml:"layout_url" json:"layout_url" bson:"layout_url"`
	// The variables below are just for information
	Description string `yaml:"description" json:"description" bson:"description"`
}

type Layouts []Layout

type SMTPServer struct {
	ID                 bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Host               string        `yaml:"host" json:"host" bson:"host"`
	Port               int           `yaml:"port" json:"port" bson:"port"`
	User               string        `yaml:"user" json:"user" bson:"user"`
	Password           string        `yaml:"password" json:"password" bson:"password"`
	From               string        `yaml:"from" json:"from" bson:"from"`
	InsecureSkipVerify bool          `yaml:"insecure_skip_verify" json:"insecure_skip_verify" bson:"insecure_skip_verify"`
}

type SMTPServers []SMTPServer

type Mail struct {
	To           []string
	Cc           []string
	Cci          []string
	Attachments  []string
	TemplateVars interface{}
	Template     MailTemplate
	SMTP         SMTPServer
}

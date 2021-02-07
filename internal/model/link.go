package model

import uuid "github.com/satori/go.uuid"

// Link is an internet hyperlink
type Link struct {
	ID           uuid.UUID `sql:"type:uuid,default:uuid_generate_v4()"`
	URI          string    `sql:",notnull"`
	Type         string
	Platform     string `sql:",notnull"`
	PersonalData bool   `sql:",notnull"`
}

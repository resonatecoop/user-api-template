package model

import uuid "github.com/satori/go.uuid"

// StreetAddress holds geographical data and marks if private
type StreetAddress struct {
	ID           uuid.UUID         `sql:"type:uuid,default:uuid_generate_v4()"`
	PersonalData bool              `sql:",notnull"`
	Data         map[string]string `pg:",hstore"`
}

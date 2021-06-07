package model

import uuid "github.com/google/uuid"

// StreetAddress holds geographical data and marks if private
type StreetAddress struct {
	ID           uuid.UUID         `bun:"type:uuid,default:uuid_generate_v4()"`
	PersonalData bool              `bun:",notnull"`
	Data         map[string]string `pg:",hstore"`
}

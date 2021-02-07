package model

import uuid "github.com/satori/go.uuid"

type GroupTaxonomy struct {
	ID   uuid.UUID `sql:"type:uuid,default:uuid_generate_v4()"`
	Type string    `sql:",notnull"`
	Name string    `sql:",notnull"`
}

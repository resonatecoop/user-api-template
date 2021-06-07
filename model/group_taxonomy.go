package model

import uuid "github.com/google/uuid"

type GroupTaxonomy struct {
	ID   uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	Type string    `bun:",notnull"`
	Name string    `bun:",notnull"`
}

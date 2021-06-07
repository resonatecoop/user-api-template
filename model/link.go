package model

import uuid "github.com/google/uuid"

// Link is an internet hyperlink
type Link struct {
	ID           uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	URI          string    `bun:",notnull"`
	Type         string
	Platform     string `bun:",notnull"`
	PersonalData bool   `bun:",notnull"`
}

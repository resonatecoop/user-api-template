package model

import "github.com/google/uuid"

// UserUploadSubmission
type UserUploadSubmission struct {
	IDRecord
	UserID uuid.UUID   `bun:"type:uuid,notnull"`
	User   *User       `bun:"rel:has-one"`
	Name   string      `bun:",notnull"`
	Files  []uuid.UUID `bun:",type:uuid[],array"`
	Active bool        `bun:"default:true,notnull"`
}

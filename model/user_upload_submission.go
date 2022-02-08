package model

import "github.com/google/uuid"

// File upload reference
type File struct {
	ID   uuid.UUID `bun:"type:uuid,notnull"`
	Name string
}

// UserUploadSubmission
type UserUploadSubmission struct {
	IDRecord
	UserID uuid.UUID `bun:"type:uuid,notnull"`
	User   *User     `bun:"rel:has-one"`
	Name   string    `bun:",notnull"`
	Files  []File    `bun:",array"`
	Active bool      `bun:"default:true,notnull"`
}

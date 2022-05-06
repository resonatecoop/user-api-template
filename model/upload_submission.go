package model

import "github.com/google/uuid"

// UploadSubmission
type UploadSubmission struct {
	IDRecord
	Active       bool        `bun:"default:true,notnull"`
	Description  string      `bun:",notnull"`
	Files        []uuid.UUID `bun:",type:uuid[],array"`
	Name         string      `bun:",notnull"`
	TrackGroup   *TrackGroup `bun:"rel:has-one"`
	TrackGroupID uuid.UUID   `bun:"type:uuid"`
	UserID       uuid.UUID   `bun:"type:uuid,notnull"`
	User         *User       `bun:"rel:has-one"`
}

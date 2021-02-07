package model

import uuid "github.com/satori/go.uuid"

// UserGroupPrivacy describes meta data for User Group private data
type UserGroupPrivacy struct {
	ID               uuid.UUID `sql:"type:uuid,default:uuid_generate_v4()"`
	Private          bool      `sql:",notnull"`
	OwnedTracks      bool      `sql:",notnull"`
	SupportedArtists bool      `sql:",notnull"`
}

package model

import (
	"time"

	"github.com/google/uuid"
)

// TrackGroup
type TrackGroup struct {
	IDRecord

	Title             string    `bun:",notnull"`
	Slug              string    `bun:",notnull"` // Slug title
	ReleaseDate       time.Time `bun:",notnull"`
	Type              string    `bun:"type:track_group_type,notnull"` // EP, LP, Single, Playlist
	Cover             []byte    `bun:",notnull"`
	DisplayArtist     string    // for display purposes, e.g. "Various" for compilation
	MultipleComposers bool      `bun:",notnull"`
	Private           bool      `bun:",notnull"`
	About             string

	OwnerId uuid.UUID `bun:"type:uuid,notnull"`
	Owner   *User     `bun:"rel:has-one"`

	UserGroupId uuid.UUID `bun:"type:uuid,default:uuid_nil()"` // track group belongs to user group, can be null if user playlist
	LabelId     uuid.UUID `bun:"type:uuid,default:uuid_nil()"`

	Tracks []uuid.UUID `bun:",type:uuid[]" pg:",array"`
	Tags   []uuid.UUID `bun:",type:uuid[]" pg:",array"`

	TerritoriesIncl []string `pg:",array"`
	CLineYear       time.Time
	PLineYear       time.Time
	CLineText       string
	PLineText       string
	RightExpiryDate time.Time
	TotalVolumes    int
	CatalogNumber   string
}

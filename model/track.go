package model

import (
	uuid "github.com/google/uuid"
)

type Track struct {
	IDRecord
	Title       string `bun:",notnull"`
	Status      string `bun:"type:track_status,notnull"`
	Enabled     bool   `bun:",notnull"`
	TrackNumber int32  `bun:",notnull"`
	Duration    float32
	Download    bool `bun:",notnull"` // Allow or disallow download

	TrackGroups     []uuid.UUID `bun:",type:uuid[],array"`
	FavoriteOfUsers []uuid.UUID `bun:",type:uuid[],array"`

	TrackServerId uuid.UUID `bun:"type:uuid,notnull"`

	OwnerId uuid.UUID `bun:"type:uuid,notnull"`
	Owner   *User     `bun:"rel:has-one"`

	UserGroupId uuid.UUID  `bun:"type:uuid,notnull"`
	UserGroup   *UserGroup `bun:"rel:has-one"` // track belongs to user group (the one who gets paid)

	Artists []uuid.UUID `bun:",type:uuid[]" pg:",array"` // for display purposes
	Tags    []uuid.UUID `bun:",type:uuid[]" pg:",array"`

	Composers  map[string]string `pg:",hstore"`
	Performers map[string]string `pg:",hstore"`

	ISRC string

	// Plays []User `pg:"many2many:plays"` Payment API
}

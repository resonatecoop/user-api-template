package model

import (
	"context"

	uuid "github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Play struct {
	IDRecord
	UserId  uuid.UUID `bun:",type:uuid,notnull"`
	TrackId uuid.UUID `bun:",type:uuid,notnull"`
	Type    string    `bun:"type:play_type,notnull"`
	Credits float32   `bun:",notnull"`
}

// Count number of times a track has been played (and paid) by a user
func CountPlays(ctx context.Context, db *bun.DB, trackId uuid.UUID, userId uuid.UUID) (int32, error) {
	play := Play{}
	count, err := db.NewSelect().
		Model(play).
		Where("user_id = ?", userId).
		Where("track_id = ?", trackId).
		Where("type = 'paid'").
		Count(ctx)
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

package model

import (
	"time"

	"github.com/google/uuid"
)

// MyGormModel mimixks GormModel but uses uuid's for ID, generated in go
type IDRecord struct {
	ID        uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time
	DeletedAt time.Time `bun:",soft_delete"`
}

// TimestampModel ...
// type TimestampModel struct {
// 	CreatedAt time.Time `bun:"default:now()"`
// 	UpdatedAt time.Time
// 	DeletedAt *time.Time
// }

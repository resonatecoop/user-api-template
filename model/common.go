package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// MyGormModel mimixks GormModel but uses uuid's for ID, generated in go
type IDRecord struct {
	ID        uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	CreatedAt time.Time `bun:"default:now()"`
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// TimestampModel ...
// type TimestampModel struct {
// 	CreatedAt time.Time `bun:"default:now()"`
// 	UpdatedAt time.Time
// 	DeletedAt *time.Time
// }

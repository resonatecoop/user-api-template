package model

import (
	"time"

	"github.com/google/uuid"
)

// Membership
type Membership struct {
	IDRecord
	UserID         uuid.UUID `bun:"type:uuid,notnull"`
	SubscriptionID string    `bun:",unique,notnull"`
	Start          time.Time
	End            time.Time
}

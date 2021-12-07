package model

import (
	"time"

	"github.com/google/uuid"
)

// UserMembership
type UserMembership struct {
	IDRecord
	UserID            uuid.UUID `bun:"type:uuid,notnull"`
	User              *User     `bun:"rel:has-one"`
	MembershipClassID uuid.UUID `bun:"type:uuid,notnull"`
	MembershipClass   *MembershipClass
	SubscriptionID    string `bun:",unique,notnull"`
	Start             time.Time
	End               time.Time
}

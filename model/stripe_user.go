package model

import "github.com/google/uuid"

// Stripe User
type StripeUser struct {
	IDRecord
	CustomerID string    `bun:",unique,notnull"`
	UserID     uuid.UUID `bun:"type:uuid,notnull"`
	User       *User     `bun:"rel:has-one"`
}
